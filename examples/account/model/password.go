package model

import scyna "github.com/scyna/core"

type Password struct {
	value string
}

func ParsePassword(password string) (Password, scyna.Error) {
	/*TODO*/
	return Password{value: password}, nil
}

func (p Password) Encode() string {
	return p.value /*TODO hash*/
}

func (p Password) String() string {
	return p.value
}
