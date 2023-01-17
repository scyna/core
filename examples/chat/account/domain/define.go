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
