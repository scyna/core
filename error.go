package scyna

import "github.com/scyna/core/internal/base"

var (
	OK                = base.OK
	SERVER_ERROR      = base.SERVER_ERROR
	BAD_REQUEST       = base.BAD_REQUEST
	PERMISSION_ERROR  = base.PERMISSION_ERROR
	REQUEST_INVALID   = base.REQUEST_INVALID
	MODULE_NOT_EXISTS = base.MODULE_NOT_EXISTS
	BAD_DATA          = base.BAD_DATA
	STREAM_ERROR      = base.STREAM_ERROR
	OBJECT_NOT_FOUND  = base.OBJECT_NOT_FOUND
	OBJECT_EXISTS     = base.OBJECT_EXISTS
)

type Error interface {
	Code() int32
	Message() string
}

func NewError(code int32, message string) Error {
	return base.NewError(code, message)
}
