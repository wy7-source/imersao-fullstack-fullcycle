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

// Na Entidade Bank, utilizamos tags para validação, json e para o ORM.
type Bank struct {
	Base     `valid:"required"` // ID, CreatedAt, UpdatedAt
	Code     string     `json:"code" gorm:"type:varchar(20)" valid:"notnull"`
	Name     string     `json:"name" gorm:"type:varchar(255)" valid:"notnull"`
	Accounts []*Account `gorm:"ForeignKey:BankID" valid:"-"` // O meu Bank, pode ter várias Accounts... A tag gorm está relacionando o nosso ID do Bank a cada conta de usuário, mas não valida notnull.
}

func (bank *Bank) isValid() error {
	// Só queremos saber se é valido ou não.
	_, err := govalidator.ValidateStruct(bank)
	if err != nil {
		return err
	}
	return nil
}

// NewBank, é o nossa função contrutora de um Bank.
func NewBank(code string, name string) (*Bank, error) {
	bank := Bank{
		Code: code,
		Name: name,
	}
	bank.ID = uuid.NewV4().String()
	bank.CreatedAt = time.Now()
	err := bank.isValid()
	if err != nil {
		return nil, err
	}
	return &bank, nil
}
