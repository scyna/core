package model

import (
	"fmt"

	scyna "github.com/scyna/core"
)

type Identity struct {
	type_  string
	number string
}

func (id Identity) Type() string {
	return id.type_
}

func (id Identity) Number() string {
	return id.number
}

func NewIdentity(type_ string, number string) (Identity, scyna.Error) {
	/*TODO: validate*/
	return Identity{type_: type_, number: number}, nil
}

func (id Identity) Marshal() string {
	return fmt.Sprintf("%s:%s", id.type_, id.number)
}

func UnmarshalIdentity(value string) (Identity, scyna.Error) {
	/*TODO*/
	return Identity{}, nil
}
