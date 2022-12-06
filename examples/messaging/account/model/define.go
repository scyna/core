package model

import (
	scyna "github.com/scyna/core"
)

const (
	CREATE_USER_URL = "/example/user/create"
	GET_USER_URL    = "/example/user/get"
)

var (
	USER_EXISTED     = &scyna.Error{Code: 100, Message: "User Existed"}
	USER_NOT_EXISTED = &scyna.Error{Code: 101, Message: "User NOT Existed"}
)
