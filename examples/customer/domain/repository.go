package domain

import scyna "github.com/scyna/core"

type IRepository interface {
	/*TODO: methods*/
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
