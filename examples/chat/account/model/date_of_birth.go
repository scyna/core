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

func (d DateOfBirth) String() string {
	return fmt.Sprintf("%d/%d/%d", d.date, d.month, d.year)
}

func (d DateOfBirth) Empty() bool {
	return d.date+d.month+d.year == 0
}

func ParseDateOfBirth(date string) (DateOfBirth, scyna.Error) {
	/*TODO*/
	return DateOfBirth{}, nil
}
