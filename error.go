package scyna

type Error interface {
	Code() int32
	Message() string
}

type errorValue struct {
	code    int32
	message string
}

func (e *errorValue) Code() int32 {
	return e.code
}

func (e *errorValue) Message() string {
	return e.Message()
}

func NewError(code int32, message string) Error {
	return &errorValue{
		code:    code,
		message: message,
	}
}
