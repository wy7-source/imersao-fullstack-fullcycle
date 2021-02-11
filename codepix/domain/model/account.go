package model

import (
	"time"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

func init() {
	// Vamos iniciar as validações assim que essa struct for iniciada.
	govalidator.SetFieldsRequiredByDefault(true)
}

// Na Entidade Account, utilizamos tags para validação, json e para o ORM.
type Account struct {
	Base      `valid:"required"`
	OwnerName string    `gorm:"column:owner_name;type:varchar(255);not null" valid:"notnull"`
	Bank      *Bank     `valid:"-"`
	BankID    string    `gorm:"column:bank_id;type:uuid;not null" valid:"-"`
	Number    string    `json:"number" gorm:"type:varchar(20)" valid:"notnull"`
	PixKeys   []*PixKey `gorm:"ForeignKey:AccountID" valid:"-"` // Na minha conta, posso ter várias PixKey's... A tag gorm está relacionando o nosso ID da Account a cada PixKey, mas não valida notnull.
}

func (account *Account) isValid() error {
	// Só queremos saber se é valido ou não.
	_, err := govalidator.ValidateStruct(account)
	if err != nil {
		return err
	}
	return nil
}

// NewAccount, é o nossa função contrutora de um Account.
func NewAccount(bank *Bank, number string, ownerName string) (*Account, error) {
	account := Account{
		Bank:      bank,
		BankID:    bank.ID,
		Number:    number,
		OwnerName: ownerName,
	}

	account.ID = uuid.NewV4().String()
	account.CreatedAt = time.Now()

	err := account.isValid()
	if err != nil {
		return nil, err
	}
	return &account, nil
}
