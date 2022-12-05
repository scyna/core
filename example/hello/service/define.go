package hello

import scyna "github.com/scyna/core"

const (
	HELLO_URL = "/example/hello/hello"
	ADD_URL   = "/example/hello/add"
)

var (
	ADD_RESULT_TOO_BIG = &scyna.Error{Code: 100, Message: "Too Big"}
)
