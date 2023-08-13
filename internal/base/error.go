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

var (
	OK                    = NewError(0, "Success")
	SERVER_ERROR          = NewError(1, "Server Error")
	BAD_REQUEST           = NewError(2, "Bad Request")
	PERMISSION_ERROR      = NewError(4, "Permission Error")
	REQUEST_INVALID       = NewError(5, "Request Invalid")
	MODULE_NOT_EXISTS     = NewError(6, "Module Not Exists")
	BAD_DATA              = NewError(7, "Bad Data")
	STREAM_ERROR          = NewError(8, "Stream Error")
	OBJECT_NOT_FOUND      = NewError(9, "Object Not Found")
	OBJECT_EXISTS         = NewError(10, "Object Exists")
	COMMAND_NOT_COMPLETED = NewError(11, "Command Not Completed")
)
