package main

import (
	"fmt"
	"sync"
)

type Observable[T any] struct {
	subs sync.Map
}

func (o *Observable[T]) Subscribe(x Observer[T]) {
	o.subs.Store(x, struct{}{})
}

func (o *Observable[T]) Unsubscribe(x Observer[T]) {
	o.subs.Delete(x)
}

func (o *Observable[T]) Fire(data T) {
	o.subs.Range(func(key, value any) bool {
		key.(Observer[T]).Notify(data)
		return true
	})
}

type Observer[T any] interface {
	Notify(data T)
}

type Person struct {
	Observable[string]
	Name string
}

func NewPerson(name string) *Person {
	return &Person{Name: name}
}

func (p *Person) CatchACold() {
	p.Fire(p.Name)
}

type DoctorService struct{}

func (d *DoctorService) Notify(data string) {
	fmt.Println("A doctor has been called for", data)
}

func main() {
	p := NewPerson("Boris")
	ds := &DoctorService{}
	p.Subscribe(ds)

	p.CatchACold()
}
