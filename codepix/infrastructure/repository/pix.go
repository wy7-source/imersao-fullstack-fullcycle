package repository

import (
	"fmt"
	"github.com/codeedu/imersao/codepix-go/domain/model"
	"github.com/jinzhu/gorm"
)
// É a nossa implementação de Repository desse nosso agregado.
// Implementando os métodos da interface PixKeyRepositoryInterface, automaticamente implementamos a interface. Sendo assim, possível a nossa Injeção de Dependência.
type PixKeyRepositoryDb struct {
	Db *gorm.DB
}

// AddBank() é o método para persistir um novo Bank.
func (r PixKeyRepositoryDb) AddBank(bank *model.Bank) error {
	err := r.Db.Create(bank).Error
	if err != nil {
		return err
	}
	return nil
}
// AddAccount() é o método para persistir uma nova Account.
func (r PixKeyRepositoryDb) AddAccount(account *model.Account) error {
	err := r.Db.Create(account).Error
	if err != nil {
		return err
	}
	return nil
}
// RegisterKey() é o método para persistir uma nova PixKey.
func (r PixKeyRepositoryDb) RegisterKey(pixKey *model.PixKey) (*model.PixKey, error) {
	err := r.Db.Create(pixKey).Error
	if err != nil {
		return nil, err
	}
	return pixKey, nil
}
// FindKeyByKind() é o método para encontrar uma PixKey por determinado Kind e ID.
func (r PixKeyRepositoryDb) FindKeyByKind(key string, kind string) (*model.PixKey, error) {
	var pixKey model.PixKey // Para injetarmos o resultado.

	// O Preload trás todo o encadeamento e suas relações além de AccountID a essa PixKey (a Account e seu Bank, e etc...). 
	r.Db.Preload("Account.Bank").First(&pixKey, "kind = ? and key = ?", kind, key) // Usamos '?' para evitar SqlInject.

	if pixKey.ID == "" {
		return nil, fmt.Errorf("no key was found")
	}
	return &pixKey, nil
}
// FindAccount() é o método para encontrar uma Account pelo seu ID.
func (r PixKeyRepositoryDb) FindAccount(id string) (*model.Account, error) {
	var account model.Account // Também para injetarmos o resultado.

	// O Preload trás todo o encadeamento e suas relações além de BankID a essa Account (o Bank e sua lista de Accounts, e etc...). 
	r.Db.Preload("Bank").First(&account, "id = ?", id) // Usamos '?' para evitar SqlInject.

	if account.ID == "" {
		return nil, fmt.Errorf("no account found")
	}
	return &account, nil
}
// FindBank() é o método para encontrar um Bank pelo seu ID.
func (r PixKeyRepositoryDb) FindBank(id string) (*model.Bank, error) {
	var bank model.Bank
	r.Db.First(&bank, "id = ?", id)

	if bank.ID == "" {
		return nil, fmt.Errorf("no bank found")
	}
	return &bank, nil
}
