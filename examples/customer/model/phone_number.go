package model

type PhoneNumber struct {
	value string
}

func (p PhoneNumber) String() string {
	return p.value
}

func (p PhoneNumber) NotProvided() bool {
	return len(p.value) == 0
}

func ParsePhoneNumber(number string) PhoneNumber {
	/*TODO: validate*/
	return PhoneNumber{value: number}
}
