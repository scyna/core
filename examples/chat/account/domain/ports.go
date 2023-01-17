package domain

import (
	scyna "github.com/scyna/core"
	"github.com/scyna/core/examples/chat/account/model"
)

type IRepository interface {
	GetAccount(email model.EmailAddress) (*model.Account, scyna.Error)
	CreateAccount(cmd *scyna.Command, account *model.Account)
}
