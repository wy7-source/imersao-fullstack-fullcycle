package model

import (
	"time"

	"github.com/asaskevich/govalidator"
)

func init() {
	// Vamos iniciar as validações assim que essa struct for iniciada.
	govalidator.SetFieldsRequiredByDefault(true)
}

// Base, é uma struct que atualmente disponibiliza para herança, campos que todas as nossas entidades tem em comum.
// Temos Tags para validação, json e para o ORM.
type Base struct {
	ID        string    `json:"id" gorm:"type:uuid;primary_key" valid:"uuid"`
	CreatedAt time.Time `json:"created_at" valid:"-"`
	UpdatedAt time.Time `json:"updated_at" valid:"-"`
}
