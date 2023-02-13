package main

import (
	"fmt"
	"strings"
)

type User struct {
	FullName string
}

func NewUser(fullName string) *User {
	return &User{FullName: fullName}
}

var allNames []string

type User2 struct {
	names []uint8 // flyweight
}

func NewUser2(fullName string) *User2 {
	getOrAdd := func(name string) uint8 {
		for i, n := range allNames {
			if name == n {
				return uint8(i)
			}
		}
		allNames = append(allNames, name)
		return uint8(len(allNames) - 1)
	}
	names := []uint8{}
	for _, s := range strings.Split(fullName, " ") {
		names = append(names, getOrAdd(s))
	}
	return &User2{names: names}
}

func (u *User2) FullName() string {
	names := []string{}
	for _, i := range u.names {
		names = append(names, allNames[i])
	}
	return strings.Join(names, " ")
}

func main() {
	john := NewUser("John Doe")
	jane := NewUser("Jane Doe")
	alsoJane := NewUser("Jane Smith")

	fmt.Println("Memory taken by users",
		len(john.FullName)+len(jane.FullName)+len(alsoJane.FullName))

	john2 := NewUser2("John Doe")
	jane2 := NewUser2("Jane Doe")
	alsoJane2 := NewUser2("Jane Smith")
	totalMem := 0
	for _, a := range allNames {
		totalMem += len(a)
	}
	totalMem += len(john2.names)
	totalMem += len(jane2.names)
	totalMem += len(alsoJane2.names)

	fmt.Println("Memory taken by users2:", totalMem)
}
