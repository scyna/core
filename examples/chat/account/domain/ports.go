package domain

import (
	scyna "github.com/scyna/core"
	"github.com/scyna/core/examples/chat/account/model"
)

type IRepository interface {
	GetAccount(email model.EmailAddress) (*model.Account, scyna.Error)
	CreateAccount(cmd *scyna.Command, account *model.Account)
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
