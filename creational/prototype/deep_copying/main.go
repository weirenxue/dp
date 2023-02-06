package main

import "fmt"

type Address struct {
	StreetAddress, City, Country string
}

type Person struct {
	Name    string
	Address *Address
}

func main() {
	john := Person{Name: "John",
		Address: &Address{
			StreetAddress: "123 London Rd",
			City:          "London",
			Country:       "UK",
		}}

	jane := john
	jane.Address = &Address{
		StreetAddress: john.Address.StreetAddress,
		City:          john.Address.City,
		Country:       john.Address.Country,
	}
	jane.Name = "Jane"
	jane.Address.StreetAddress = "321 Baker Rd"

	fmt.Println(john, john.Address)
	fmt.Println(jane, jane.Address)
}
