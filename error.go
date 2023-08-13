package scyna

import "github.com/scyna/core/internal/base"

var (
	OK                 = NewError(0, "Success")
	SERVER_ERROR       = NewError(1, "Server Error")
	BAD_REQUEST        = NewError(2, "Bad Request")
	PERMISSION_ERROR   = NewError(4, "Permission Error")
	REQUEST_INVALID    = NewError(5, "Request Invalid")
	MODULE_NOT_EXISTED = NewError(6, "Module Not Existed")
	BAD_DATA           = NewError(7, "Bad Data")
	STREAM_ERROR       = NewError(8, "Stream Error")
)

type Error interface {
	Code() int32
	Message() string
}

func NewError(code int32, message string) Error {
	return base.NewError(code, message)
}
