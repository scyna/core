package scyna

import "fmt"

type Error interface {
	Code() string
	Message() string
}

type errorData struct {
	code    string
	message string
}

func (e *errorData) Code() string {
	return e.code
}

func (e *errorData) Error() string {
	return fmt.Sprintf("Code:%s, Message: %s", e.code, e.message)
}

func (e *errorData) Message() string {
	return e.message
}

func NewError(code string, message string) Error {
	return &errorData{
		code:    code,
		message: message,
	}
}

var (
	OK                    = NewError("OK", "Success")
	SERVER_ERROR          = NewError("ServerError", "Server Error")
	BAD_REQUEST           = NewError("BadRequest", "Bad Request")
	PERMISSION_ERROR      = NewError("PermissionError", "Permission Error")
	REQUEST_INVALID       = NewError("RequestInvalid", "Request Invalid")
	MODULE_NOT_EXISTS     = NewError("ModuleNotExists", "Module Not Exists")
	BAD_DATA              = NewError("BadData", "Bad Data")
	STREAM_ERROR          = NewError("StreamError", "Stream Error")
	OBJECT_NOT_FOUND      = NewError("ObjectNotFound", "Object Not Found")
	OBJECT_EXISTS         = NewError("ObjectExists", "Object Exists")
	COMMAND_NOT_COMPLETED = NewError("CommandNotCompleted", "Command Not Completed")
)
