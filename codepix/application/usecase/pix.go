package usecase

import (
	"github.com/codeedu/imersao/codepix-go/domain/model"
)

// Injetamos quem já implementa a interface.
type PixUseCase struct {
	PixKeyRepository model.PixKeyRepositoryInterface
}
// RegisterKey é o método que Cria uma nova PixKey.
func (p *PixUseCase) RegisterKey(key string, kind string, accountId string) (*model.PixKey, error) {
	account, err := p.PixKeyRepository.FindAccount(accountId)
	if err != nil {
		return nil, err
	}

	pixKey, err := model.NewPixKey(kind, account, key)
	if err != nil {
		return nil, err
	}

	p.PixKeyRepository.RegisterKey(pixKey)
	if pixKey.ID == "" {
		return nil, err
	}

	return pixKey, nil
}
// FindKey é o método que Busca a PixKey pelo seu ID e Kind.
func (p *PixUseCase) FindKey(key string, kind string) (*model.PixKey, error) {
	pixKey, err := p.PixKeyRepository.FindKeyByKind(key, kind)
	if err != nil {
		return nil, err
	}
	return pixKey, nil
}