package model

import scyna "github.com/scyna/core"

type EmailAddress struct {
	addr string
}

func ParseEmail(addr string) (EmailAddress, scyna.Error) {
	/*TODO: validate addr*/
	return EmailAddress{addr: addr}, nil
}

func (e EmailAddress) ToString() string {
	return e.addr
}
