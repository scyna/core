package base

import "fmt"

type Error struct {
	code    int32
	message string
}

func (e *Error) Code() int32 {
	return e.code
}

func (e *Error) Error() string {
	return fmt.Sprintf("Code:%d, Message: %s", e.code, e.message)
}

func (e *Error) Message() string {
	return e.message
}

func NewError(code int32, message string) *Error {
	return &Error{
		code:    code,
		message: message,
	}
}
