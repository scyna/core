package model

import scyna "github.com/scyna/core"

type CustomerName struct {
	value string
}

func (n CustomerName) String() string {
	return n.value
}

func NewCustomerName(name string) (CustomerName, scyna.Error) {
	/*TODO*/
	return CustomerName{}, nil
}
