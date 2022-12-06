package model

import (
	scyna "github.com/scyna/core"
)

const (
	CREATE_USER_URL         = "/chat/account/create"
	GET_USER_URL            = "/chat/account/get"
	ACCOUNT_CREATED_CHANNEL = "chat.account.user_created"
)

var (
	USER_EXISTED     = &scyna.Error{Code: 100, Message: "User Existed"}
	USER_NOT_EXISTED = &scyna.Error{Code: 101, Message: "User NOT Existed"}
)
