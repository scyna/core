package domain

import (
	scyna "github.com/scyna/core"
	"github.com/scyna/core/examples/account/model"
)

type IRepository interface {
	GetAccountByEmail(email model.EmailAddress) (*model.Account, scyna.Error)
	GetAccountByID(ID uint64) (*model.Account, scyna.Error)
	LoadPassword(acc *model.Account) scyna.Error

	CreateAccount(cmd *scyna.Command, account *model.Account) scyna.Error
	UpdatePassword(cmd *scyna.Command, account *model.Account) scyna.Error
}

type RepositoryCreator func(LOG scyna.Logger) IRepository

var repositoryCreator RepositoryCreator

func LoadRepository(LOG scyna.Logger) IRepository {
	if repositoryCreator == nil {
		panic("No RepositoryCreator attached")
	}
	return repositoryCreator(LOG)
}

func AttachRepositoryCreator(rc RepositoryCreator) {
	repositoryCreator = rc
}
