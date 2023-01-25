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

func (a StreetAddress) IsEmpty() bool {
	return len(a.street) == 0
}

type streetAddressBuilder struct {
	street   string
	district string
	province string
	country  string
}

func NewStreetAddress() *streetAddressBuilder {
	return &streetAddressBuilder{}
}

func (b *streetAddressBuilder) Modify(address StreetAddress) *streetAddressBuilder {
	b.street = address.street
	b.district = address.district
	b.province = address.province
	b.country = address.country
	return b
}

func (b *streetAddressBuilder) WithStreet(street string) *streetAddressBuilder {
	b.street = street
	return b
}

func (b *streetAddressBuilder) WithDistrict(district string) *streetAddressBuilder {
	b.district = district
	return b
}

func (b *streetAddressBuilder) WithProvive(province string) *streetAddressBuilder {
	b.province = province
	return b
}

func (b *streetAddressBuilder) WithCountry(country string) *streetAddressBuilder {
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
