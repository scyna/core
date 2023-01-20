package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	scyna "github.com/scyna/core"
)

type EmailAddress struct {
	addr string
}

func ParseEmail(addr string) (EmailAddress, scyna.Error) {
	if validation.Validate(addr, validation.Required, is.Email) != nil {
		return EmailAddress{}, BAD_EMAIL
	}
	return EmailAddress{addr: addr}, nil
}

func (e EmailAddress) ToString() string {
	return e.addr
}
