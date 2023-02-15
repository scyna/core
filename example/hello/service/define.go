package service

import scyna "github.com/scyna/core"

const (
	HELLO_URL = "/ex/hello/hello"
	ADD_URL   = "/ex/hello/add"
)

var (
	ADD_RESULT_TOO_BIG = scyna.NewError(100, "Too Big")
)
