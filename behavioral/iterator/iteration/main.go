package main

import "fmt"

type Person struct {
	FirstName, MiddleName, LastName string
}

func (p *Person) Names() [3]string {
	return [3]string{p.FirstName, p.MiddleName, p.LastName}
}

func (p *Person) NamesGenerator() <-chan string {
	out := make(chan string)
	go func() {
		out <- p.FirstName
		if len(p.MiddleName) > 0 {
			out <- p.MiddleName
		}
		out <- p.LastName
		close(out)
	}()
	return out
}

type PersonNameIterator struct {
	person  *Person
	current int
}

func NewPersonNameIterator(p *Person) *PersonNameIterator {
	return &PersonNameIterator{person: p, current: -1}
}

func (p *PersonNameIterator) MoveNext() bool {
	p.current++
	return p.current < 3
}
func (p *PersonNameIterator) Value() string {
	switch p.current {
	case 0:
		return p.person.FirstName
	case 1:
		return p.person.MiddleName
	case 2:
		return p.person.LastName
	}
	panic("We should not be here!")
}

func main() {
	fmt.Println("# p.Names()")
	p := Person{"Alexander", "Graham", "Bell"}
	for _, n := range p.Names() {
		fmt.Println(n)
	}

	fmt.Println("\n# p.NamesGenerator()")
	p = Person{"Alexander", "", "Bell"}
	for n := range p.NamesGenerator() {
		fmt.Println(n)
	}

	fmt.Println("\n# PersonNameIterator")
	p = Person{"Alexander", "Graham", "Bell"}
	for it := NewPersonNameIterator(&p); it.MoveNext(); {
		fmt.Println(it.Value())
	}
}
