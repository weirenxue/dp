package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

type Address struct {
	Suite               int
	StreetAddress, City string
}

type Employee struct {
	Name   string
	Office Address
}

func (p *Employee) DeepCopy() *Employee {
	b := bytes.Buffer{}
	e := gob.NewEncoder(&b)
	_ = e.Encode(p)

	result := Employee{}
	d := gob.NewDecoder(&b)
	d.Decode(&result)
	return &result
}

var mainOffice = Employee{Name: "", Office: Address{Suite: 0, StreetAddress: "123 London Rd", City: "London"}}
var auxOffice = Employee{Name: "", Office: Address{Suite: 0, StreetAddress: "321 Baker St", City: "London"}}

func newEmployee(proto *Employee, name string, suite int) *Employee {
	p := proto.DeepCopy()
	p.Name = name
	p.Office.Suite = suite
	return p
}

func NewMainOfficeEmployee(name string, suite int) *Employee {
	return newEmployee(&mainOffice, name, suite)
}
func NewAuxOfficeEmployee(name string, suite int) *Employee {
	return newEmployee(&auxOffice, name, suite)
}

func main() {
	john := NewMainOfficeEmployee("John", 100)
	jane := NewAuxOfficeEmployee("Jane", 200)

	fmt.Println(john)
	fmt.Println(jane)
}
