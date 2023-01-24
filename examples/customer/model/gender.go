package model

import scyna "github.com/scyna/core"

type Gender struct {
	value string
}

func (g Gender) String() string {
	return g.value
}

func (g Gender) Empty() bool {
	return len(g.value) == 0
}

func NewGender(gender string) (Gender, scyna.Error) {
	/*TODO: validate*/
	return Gender{value: gender}, nil
}
