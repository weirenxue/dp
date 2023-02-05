package main

import "fmt"

type Person struct {
	// address
	StreetAddress, Postcode, City string
	// job
	CompanyName, Position string
	AnnualIncome          int
}

type PersonBuilder struct {
	person *Person
}

func NewPersonBuilder() *PersonBuilder {
	return &PersonBuilder{person: &Person{}}
}

func (p *PersonBuilder) Build() *Person {
	return p.person
}

func (p *PersonBuilder) Lives() *PersonAddressBuilder {
	return &PersonAddressBuilder{PersonBuilder: *p}
}

func (p *PersonBuilder) Works() *PersonJobBuilder {
	return &PersonJobBuilder{PersonBuilder: *p}
}

type PersonAddressBuilder struct {
	PersonBuilder
}

func (p *PersonAddressBuilder) At(stressAddress string) *PersonAddressBuilder {
	p.person.StreetAddress = stressAddress
	return p
}

func (p *PersonAddressBuilder) In(city string) *PersonAddressBuilder {
	p.person.City = city
	return p
}

func (p *PersonAddressBuilder) WithPostcode(postcode string) *PersonAddressBuilder {
	p.person.Postcode = postcode
	return p
}

type PersonJobBuilder struct {
	PersonBuilder
}

func (p *PersonJobBuilder) At(companyName string) *PersonJobBuilder {
	p.person.CompanyName = companyName
	return p
}

func (p *PersonJobBuilder) AsA(position string) *PersonJobBuilder {
	p.person.Position = position
	return p
}

func (p *PersonJobBuilder) Earning(annualIncome int) *PersonJobBuilder {
	p.person.AnnualIncome = annualIncome
	return p
}

func main() {
	pb := NewPersonBuilder()
	pb.
		Lives().At("123 London Road").In("London").WithPostcode("SW12BC").
		Works().At("Fabrikam").AsA("Programmer").Earning(123000)
	p := pb.Build()
	fmt.Println(*p)
}
