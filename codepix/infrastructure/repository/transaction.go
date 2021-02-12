package repository

import (
	"fmt"
	"github.com/codeedu/imersao/codepix-go/domain/model"
	"github.com/jinzhu/gorm"
)
// É a nossa implementação de Repository desse nosso agregado.
// Implementando os métodos da interface TransactionRepositoryInterface, automaticamente implementamos a interface. Sendo assim, possível a nossa Injeção de Dependência.
type TransactionRepositoryDb struct {
	Db *gorm.DB
}
// Register() é o método para persistir uma nova Transaction.
func (t *TransactionRepositoryDb) Register(transaction *model.Transaction) error {
	err := t.Db.Create(transaction).Error
	if err != nil {
		return err
	}
	return nil
}
// Save() é o método para atualizar uma Transaction.
func (t *TransactionRepositoryDb) Save(transaction *model.Transaction) error {
	err := t.Db.Save(transaction).Error
	if err != nil {
		return err
	}
	return nil
}
// Find() é o método para encontrar uma Transaction pelo ID.
func (t *TransactionRepositoryDb) Find(id string) (*model.Transaction, error) {
	var transaction model.Transaction // Para injetarmos o resultado.

	// O Preload trás todo o encadeamento e suas relações além do AccountID, dessa Transaction (a Account e seu Bank, e etc...). 
	t.Db.Preload("AccountFrom.Bank").First(&transaction, "id = ?", id)

	if transaction.ID == "" {
		return nil, fmt.Errorf("no key was found")
	}
	return &transaction, nil
}
