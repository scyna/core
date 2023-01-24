package model

import (
	"fmt"

	scyna "github.com/scyna/core"
)

type DateOfBirth struct {
	date  int
	month int
	year  int
}

func (dob DateOfBirth) String() string {
	return fmt.Sprintf("%d/%d/%d", dob.date, dob.month, dob.year)
}

func ParseDateOfBirth(date string) (DateOfBirth, scyna.Error) {
	/*TODO*/
	return DateOfBirth{}, nil
}
