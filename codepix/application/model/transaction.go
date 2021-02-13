package model

import (
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
)
// Transaction é uma Dto para parse de dados vindos do Kafka.
type Transaction struct {
	ID           string  `json:"id" validate:"required,uuid4"`
	AccountID    string  `json:"accountId" validate:"required,uuid4"`
	Amount       float64 `json:"amount" validate:"required,numeric"`
	PixKeyTo     string  `json:"pixKeyTo" validate:"required"`
	PixKeyKindTo string  `json:"pixKeyKindTo" validate:"required"`
	Description  string  `json:"description" validate:"required"`
	Status       string  `json:"status" validate:"-"`
	Error        string  `json:"error"`
}
// isValid é o método para validar manualmente a Dto.
func (t *Transaction) isValid() error {
	v := validator.New()
	err := v.Struct(t)
	if err != nil {
		fmt.Errorf("Error during Transaction validation: %s", err.Error())
		return err
	}
	return nil
}
// ParseJson é o método para deserealizar do Kafka para a nossa Dto.
func (t *Transaction) ParseJson(data []byte) error {
	err := json.Unmarshal(data, t)
	if err != nil {
		return err
	}

	err = t.isValid()
	if err != nil {
		return err
	}

	return nil
}
// ToJson é o método para serealizar a nossa Dto para enviar pelo Kafka.
func (t *Transaction) ToJson() ([]byte, error) {
	err := t.isValid()
	if err != nil {
		return nil, err
	}

	result, err := json.Marshal(t)
	if err != nil {
		return nil, nil
	}

	return result, nil
}
// NewTransaction é o construtor da Dto.
func NewTransaction() *Transaction {
	return &Transaction{}
}
