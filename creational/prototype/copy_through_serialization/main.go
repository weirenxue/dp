package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

type Address struct {
	StreetAddress, City, Country string
}

type Person struct {
	Name    string
	Address *Address
	Friends []string
}

func (p *Person) DeepCopy() *Person {
	b := bytes.Buffer{}
	e := gob.NewEncoder(&b)
	_ = e.Encode(p)

	fmt.Println(b.String())

	q := Person{}
	d := gob.NewDecoder(&b)
	d.Decode(&q)
	return &q
}

func main() {
	john := &Person{Name: "John",
		Address: &Address{
			StreetAddress: "123 London Rd",
			City:          "London",
			Country:       "UK",
		},
		Friends: []string{"Chris", "Matt"}}

	jane := john.DeepCopy()
	jane.Name = "Jane"
	jane.Address.StreetAddress = "321 Baker Rd"
	jane.Friends = append(jane.Friends, "Angela")

	john.Friends = append(john.Friends, "Jane")

	fmt.Println(john, john.Address)
	fmt.Println(jane, jane.Address)
}
