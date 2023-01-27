package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	scyna "github.com/scyna/core"
)

type EmailAddress struct {
	value string
}

func ParseEmail(email string) (EmailAddress, scyna.Error) {
	if validation.Validate(email, validation.Required, is.Email) != nil {
		return EmailAddress{}, BAD_EMAIL
	}
	return EmailAddress{value: email}, nil
}

func (e EmailAddress) String() string {
	return e.value
}
