package model

import scyna "github.com/scyna/core"

type PhoneNumber struct {
	number string
}

type Date struct {
	date string
}

type Gender struct {
	gender string
}

type Password struct {
	password string
}

func ParseGender(gender string) (Gender, scyna.Error) {
	/*TODO*/
	return Gender{gender: gender}, nil
}

func ParsePhoneNumber(number string) (PhoneNumber, scyna.Error) {
	/*TODO: validate number*/
	return PhoneNumber{number: number}, nil
}

func ParseDate(date string) (Date, scyna.Error) {
	/*TODO*/
	return Date{date: date}, nil
}

func ParsePassword(password string) (Password, scyna.Error) {
	/*TODO*/
	return Password{password: password}, nil
}
