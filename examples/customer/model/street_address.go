package model

type StreetAddress struct {
	street   string
	district string
	province string
	country  string
}

func (a StreetAddress) Street() string {
	return a.street
}

func (a StreetAddress) District() string {
	return a.district
}

func (a StreetAddress) Province() string {
	return a.province
}

func (a StreetAddress) Country() string {
	return a.country
}
