package model

import (
	scyna "github.com/scyna/core"
)

var (
	USER_EXISTED     = scyna.NewError(100, "User Existed")
	USER_NOT_EXISTED = scyna.NewError(101, "User NOT Existed")
	BAD_EMAIL        = scyna.NewError(102, "Bad Email")
)
