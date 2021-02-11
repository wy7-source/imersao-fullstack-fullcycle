package model

import (
	"errors"
	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
	"time"
)

// Essas enumerações são os possiveis status da nossa transaction.
const (
	TransactionPending   string = "pending"
	TransactionCompleted string = "completed"
	TransactionError     string = "error"
	TransactionConfirmed string = "confirmed"
)
// Nosso agregado Transaction, consiste em: Se eu tenho uma Transaction, logo eu tenho uma PixKey e uma Account, que tem um Bank.
// TransactionRepositoryInterface é a interface para injeção de dependencia no repository do nosso Agregado.
type TransactionRepositoryInterface interface {
	Register(transaction *Transaction) error
	Save(transaction *Transaction) error
	Find(id string) (*Transaction, error)
}

type Transactions struct {
	Transaction []Transaction
}

// Na Entidade Transaction, utilizamos tags para validação, json e para o ORM.
type Transaction struct {
	Base              `valid:"required"`
	AccountFrom       *Account `valid:"-"`
	AccountFromID     string   `gorm:"column:account_from_id;type:uuid;" valid:"notnull"`
	Amount            float64  `json:"amount" gorm:"type:float" valid:"notnull"`
	PixKeyTo          *PixKey  `valid:"-"`
	PixKeyIdTo        string   `gorm:"column:pix_key_id_to;type:uuid;" valid:"notnull"`
	Status            string   `json:"status" gorm:"type:varchar(20)" valid:"notnull"`
	Description       string   `json:"description" gorm:"type:varchar(255)" valid:"-"`
	CancelDescription string   `json:"cancel_description" gorm:"type:varchar(255)" valid:"-"`
}

func init() {
	// Vamos iniciar as validações assim que essa struct for iniciada.
	govalidator.SetFieldsRequiredByDefault(true)
}

func (t *Transaction) isValid() error {
	// Só queremos saber se é valido ou não.
	_, err := govalidator.ValidateStruct(t)

	// RN: Não pode transferir menos que 0.
	if t.Amount <= 0 {
		return errors.New("the amount must be greater than 0")
	}

	// RN: Status não pode ser diferente dos status possíveis.
	if t.Status != TransactionPending && t.Status != TransactionCompleted && t.Status != TransactionError {
		return errors.New("invalid status for the transaction")
	}

	// RN: Não pode transferir para sí mesmo.
	if t.PixKeyTo.AccountID == t.AccountFromID {
		return errors.New("the source and destination account cannot be the same")
	}

	if err != nil {
		return err
	}
	return nil
}

// Abaixo, seguindo a linguagem universal do negócio, são as unicas operações que essa entidade faz.
// Complete é a operação que completa a Transaction.
func (t *Transaction) Complete() error {
	t.Status = TransactionCompleted
	t.UpdatedAt = time.Now()
	err := t.isValid()
	return err
}

// INFO: Há a possibilidade de ter uma operação de Confirmação da Transação, mas seria basicamente uma cópia da Complete(). Então optamos por deixar isso no proprio usecase.

// Cancel é a operação que cancela a Transaction.
func (t *Transaction) Cancel(description string) error {
	t.Status = TransactionError
	t.CancelDescription = description
	t.UpdatedAt = time.Now()
	err := t.isValid()
	return err
}

// NewTransaction, é o nossa função contrutora de um PixKey.
func NewTransaction(accountFrom *Account, amount float64, pixKeyTo *PixKey, description string, id string) (*Transaction, error) {
	transaction := Transaction{
		AccountFrom:   accountFrom,
		AccountFromID: accountFrom.ID,
		Amount:        amount,
		PixKeyTo:      pixKeyTo,
		PixKeyIdTo:    pixKeyTo.ID,
		Status:        TransactionPending, // Uma transação sempre inicia com status Pending.
		Description:   description,
	}
	if id == "" {
		transaction.ID = uuid.NewV4().String()
	} else {
		transaction.ID = id
	}
	transaction.CreatedAt = time.Now()
	err := transaction.isValid()
	if err != nil {
		return nil, err
	}
	return &transaction, nil
}
