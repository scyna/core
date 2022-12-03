package user

import (
	scyna "github.com/scyna/core"
)

var (
	USER_EXISTED     = &scyna.Error{Code: 100, Message: "User Existed"}
	USER_NOT_EXISTED = &scyna.Error{Code: 101, Message: "User NOT Existed"}
)
