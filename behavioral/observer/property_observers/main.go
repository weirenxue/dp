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
	p.age = age
	p.Fire(PropertyChange{Name: "age", Value: p.age})
}

type TrafficManagement struct {
	o *Observable
}

func (t *TrafficManagement) Notify(data any) {
	if pc, ok := data.(PropertyChange); ok {
		if pc.Value.(int) >= 16 {
			fmt.Println("Congrats, you can drive")
			t.o.Unsubscribe(t)
		}
	}
}

func main() {
	p := NewPerson(15)
	tm := &TrafficManagement{&p.Observable}
	p.Subscribe(tm)

	for i := 16; i <= 20; i++ {
		fmt.Println("Setting age to", i)
		p.SetAge(i)
	}
}
