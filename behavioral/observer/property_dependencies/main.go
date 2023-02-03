package main

import (
	"fmt"
	"sync"
)

type Observable struct {
	subs sync.Map
}

func (o *Observable) Subscribe(x Observer) {
	o.subs.Store(x, struct{}{})
}

func (o *Observable) Unsubscribe(x Observer) {
	o.subs.Delete(x)
}

func (o *Observable) Fire(data PropertyChange) {
	o.subs.Range(func(key, value any) bool {
		key.(Observer).Notify(data)
		return true
	})
}

type PropertyChange struct {
	Name  string
	Value any
}

type Observer interface {
	Notify(data any)
}

type Person struct {
	Observable
	age int
}

func NewPerson(age int) *Person {
	return &Person{age: age}
}

func (p *Person) Age() int {
	return p.age
}

func (p *Person) SetAge(age int) {
	if age == p.age {
		return
	}

	canVoteCache := p.CanVote()

	p.age = age
	p.Fire(PropertyChange{Name: "age", Value: p.age})

	if canVoteCache != p.CanVote() {
		p.Fire(PropertyChange{Name: "CanVote", Value: p.CanVote()})
	}
}

func (p *Person) CanVote() bool {
	return p.age >= 18
}

type ElectroalRoll struct{}

func (e *ElectroalRoll) Notify(data any) {
	if pc, ok := data.(PropertyChange); ok {
		if pc.Name == "CanVote" && pc.Value.(bool) {
			fmt.Println("Congratulations, you can vote!")
		}
	}
}

func main() {
	p := NewPerson(0)
	er := &ElectroalRoll{}
	p.Subscribe(er)

	for i := 15; i < 25; i++ {
		fmt.Println("Setting age to", i)
		p.SetAge(i)
	}
}
