package main

import "fmt"

type Person struct {
	Name    string
	Room    *ChatRoom
	chatLog []string
}

func NewPerson(name string) *Person {
	return &Person{Name: name}
}

func (p *Person) Receive(src, msg string) {
	s := fmt.Sprintf("%s: '%s'\n", src, msg)
	fmt.Printf("[%s's chat session] %s", p.Name, s)
	p.chatLog = append(p.chatLog, s)
}

func (p *Person) Say(msg string) {
	p.Room.Broadcast(p.Name, msg)
}

func (p *Person) PrivateMessage(dst, msg string) {
	p.Room.Message(p.Name, dst, msg)
}

type ChatRoom struct {
	people []*Person
}

func (c *ChatRoom) Broadcast(src, msg string) {
	for _, p := range c.people {
		if p.Name != src {
			p.Receive(src, msg)
		}
	}
}

func (c *ChatRoom) Join(p *Person) {
	s := fmt.Sprintf("%s joins the chat", p.Name)
	c.Broadcast("Room", s)

	c.people = append(c.people, p)
	p.Room = c
}

func (c *ChatRoom) Message(src, dst, msg string) {
	for _, p := range c.people {
		if p.Name == dst {
			p.Receive(src, msg)
		}
	}
}

func main() {
	room := &ChatRoom{}

	john := NewPerson("John")
	jane := NewPerson("Jane")

	room.Join(john)
	room.Join(jane)
	john.Say("hi room")
	jane.Say("oh, hey john")

	simon := NewPerson("Simon")
	room.Join(simon)
	simon.Say("hi everyone!")

	jane.PrivateMessage(simon.Name, "glad you could join us!")
}
