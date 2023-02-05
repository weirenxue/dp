package main

import (
	"fmt"
	"strings"
)

type email struct {
	from, to, subject, body string
}

type EmailBuilder struct {
	email email
}

func (e *EmailBuilder) From(from string) *EmailBuilder {
	if !strings.Contains(from, "@") {
		panic("email should contain @")
	}
	e.email.from = from
	return e
}

func (e *EmailBuilder) To(to string) *EmailBuilder {
	e.email.to = to
	return e
}

func (e *EmailBuilder) Subject(subject string) *EmailBuilder {
	e.email.subject = subject
	return e
}

func (e *EmailBuilder) Body(body string) *EmailBuilder {
	e.email.body = body
	return e
}

func sendEmailImpl(email *email) {
	fmt.Println(*email)
}

type build func(b *EmailBuilder)

func SendEmail(build build) {
	b := EmailBuilder{}
	build(&b)
	sendEmailImpl(&b.email)
}

func main() {
	SendEmail(func(b *EmailBuilder) {
		b.From("foo@bar.com").
			To("bar@baz.com").
			Subject("Meeting").
			Body("Do you want to meet?")
	})
}
