package model

import "fmt"

type Name struct {
	value string
}

func NewName(name string) (Name, error) {
	if len(name) > 128 || name == "" {
		return Name{}, fmt.Errorf("invalid name, must be withn 20 charcters and non-empty")
	}

	return Name{value: name}, nil
}

func (n Name) String() string {
	return n.value
}

func (n Name) MarshalText() ([]byte, error) {
	return []byte(n.value), nil
}

func (n *Name) UnmarshalText(d []byte) error {
	var err error
	*n, err = NewName(string(d))
	return err
}
