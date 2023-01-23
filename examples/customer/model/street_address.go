package model

import scyna "github.com/scyna/core"

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

type streetAddressBuilder struct {
	street   string
	district string
	province string
	country  string
}

func NewStreetAddressBuilder() *streetAddressBuilder {
	return &streetAddressBuilder{}
}

func (b *streetAddressBuilder) SetStreet(street string) *streetAddressBuilder {
	b.street = street
	return b
}

func (b *streetAddressBuilder) SetDistrict(district string) *streetAddressBuilder {
	b.district = district
	return b
}

func (b *streetAddressBuilder) SetProvive(province string) *streetAddressBuilder {
	b.province = province
	return b
}

func (b *streetAddressBuilder) SetCountry(country string) *streetAddressBuilder {
	b.country = country
	return b
}

func (b *streetAddressBuilder) Build() (StreetAddress, scyna.Error) {
	/*TODO validate*/
	return StreetAddress{
		street:   b.street,
		district: b.district,
		province: b.province,
		country:  b.country,
	}, nil
}
