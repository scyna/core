package model

import scyna "github.com/scyna/core"

type PhoneNumber struct {
	value string
}

func ParsePhoneNumber(number string) (PhoneNumber, scyna.Error) {
	/*TODO: validate number*/
	return PhoneNumber{value: number}, nil
}

func (p PhoneNumber) String() string {
	return p.value
}

func (p PhoneNumber) NotProvided() bool {
	return len(p.value) == 0
}
