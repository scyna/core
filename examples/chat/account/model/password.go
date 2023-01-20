package model

import scyna "github.com/scyna/core"

type Password struct {
	password string
}

func ParsePassword(password string) (Password, scyna.Error) {
	/*TODO*/
	return Password{password: password}, nil
}

func (p Password) Encode() string {
	return p.password
}
