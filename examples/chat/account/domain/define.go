package domain

import (
	scyna "github.com/scyna/core"
)

const (
	CREATE_USER_URL         = "/chat/account/create"
	GET_USER_URL            = "/chat/account/get"
	ACCOUNT_CREATED_CHANNEL = "chat.account.user_created"
)

var (
	USER_EXISTED     = scyna.NewError(100, "User Existed")
	USER_NOT_EXISTED = scyna.NewError(101, "User NOT Existed")
	BAD_EMAIL        = scyna.NewError(102, "Bad Email")
)

type RepositoryCreator func(LOG scyna.Logger) IRepository

var repositoryCreator RepositoryCreator

func LoadAccountRepository(LOG scyna.Logger) IRepository {
	if repositoryCreator == nil {
		panic("No RepositoryCreator attached")
	}
	return repositoryCreator(LOG)
}

func AttachRepositoryCreator(rc RepositoryCreator) {
	repositoryCreator = rc
}
