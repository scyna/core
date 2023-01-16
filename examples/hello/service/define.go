package hello

import scyna "github.com/scyna/core"

const (
	HELLO_URL = "/example/hello/hello"
	ADD_URL   = "/example/hello/add"
)

var (
	ADD_RESULT_TOO_BIG = scyna.NewError(100, "Too Big")
)
