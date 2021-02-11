package model

import (
	"errors"
	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
	"time"
)

// Nosso agregado PixKey, consiste em: Se eu tenho uma PixKey, logo eu tenho uma Account, que tem um Bank.
// PixKeyRepositoryInterface é a interface para injeção de dependencia no repository do nosso Agregado.
type PixKeyRepositoryInterface interface {
	RegisterKey(pixKey *PixKey) (*PixKey, error)
	FindKeyByKind(key string, kind string) (*PixKey, error)
	AddBank(bank *Bank) error
	AddAccount(account *Account) error
	FindAccount(id string) (*Account, error)
}

func init() {
	// Vamos iniciar as validações assim que essa struct for iniciada.
	govalidator.SetFieldsRequiredByDefault(true)
}

// Na Entidade PixKey, utilizamos tags para validação, json e para o ORM.
type PixKey struct {
	Base      `valid:"required"`
	Kind      string   `json:"kind" gorm:"type:varchar(20)" valid:"notnull"`
	Key       string   `json:"key" gorm:"type:varchar(255)" valid:"notnull"`
	AccountID string   `gorm:"column:account_id;type:uuid;not null" valid:"-"`
	Account   *Account `valid:"-"`
	Status    string   `json:"status" gorm:"type:varchar(20)" valid:"notnull"`
}

func (p *PixKey) isValid() error {
	// Só queremos saber se é valido ou não.
	_, err := govalidator.ValidateStruct(p)

	// Validações simples, mas existem.
	if p.Kind != "email" && p.Kind != "cpf" {
		return errors.New("invalid type of key")
	}

	if p.Status != "active" && p.Status != "inactive" {
		return errors.New("invalid status")
	}

	if err != nil {
		return err
	}
	return nil
}

// NewPixKey, é o nossa função contrutora de um PixKey.
func NewPixKey(kind string, account *Account, key string) (*PixKey, error) {
	pixKey := PixKey{
		Kind:      kind,
		Key:       key,
		Account:   account,
		AccountID: account.ID,
		Status:    "active",
	}
	pixKey.ID = uuid.NewV4().String()
	pixKey.CreatedAt = time.Now()
	err := pixKey.isValid()
	if err != nil {
		return nil, err
	}
	return &pixKey, nil
}
