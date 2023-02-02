package main

import "fmt"

var overdraftLimit = -500

type BankAccount struct {
	balance int
}

func (b *BankAccount) Deposit(amount int) {
	b.balance += amount
	fmt.Println("Deposit", amount, "\b, balance is now", b.balance)
}

func (b *BankAccount) Withdraw(amount int) bool {
	if b.balance-amount > overdraftLimit {
		b.balance -= amount
		fmt.Println("Withdraw", amount, "\b, balance is now", b.balance)
		return true
	}
	return false
}

// Define the command interface.
type Command interface {
	Call()
	Undo()
	Succeeded() bool
	SetSucceeded(value bool)
}

// Define the action of the command.
type Action int

const (
	Deposit Action = iota
	Withdraw
)

// Implement the command interface.
// Make the command an object.
type BankAccountCommand struct {
	account   *BankAccount
	action    Action
	amount    int
	succeeded bool
}

func NewBankAccountCommand(account *BankAccount, action Action, amount int) *BankAccountCommand {
	return &BankAccountCommand{account: account, action: action, amount: amount}
}

func (b *BankAccountCommand) Call() {
	switch b.action {
	case Deposit:
		b.account.Deposit(b.amount)
		b.succeeded = true
	case Withdraw:
		b.succeeded = b.account.Withdraw(b.amount)
	}
}

// Undo allows us to undo the command if the command was successful.
func (b *BankAccountCommand) Undo() {
	if !b.succeeded {
		return
	}
	switch b.action {
	case Deposit:
		b.account.Withdraw(b.amount)
	case Withdraw:
		b.account.Deposit(b.amount)
	}
}

func (b *BankAccountCommand) Succeeded() bool {
	return b.succeeded
}

func (b *BankAccountCommand) SetSucceeded(value bool) {
	b.succeeded = value
}

type CompositeBankAccountCommand struct {
	commands []Command
}

func (c *CompositeBankAccountCommand) Call() {
	for _, cmd := range c.commands {
		cmd.Call()
	}
}

func (c *CompositeBankAccountCommand) Undo() {
	for idx := range c.commands {
		c.commands[len(c.commands)-idx-1].Undo()
	}
}

func (c *CompositeBankAccountCommand) Succeeded() bool {
	for _, cmd := range c.commands {
		if !cmd.Succeeded() {
			return false
		}
	}
	return true
}

func (c *CompositeBankAccountCommand) SetSucceeded(value bool) {
	for _, cmd := range c.commands {
		cmd.SetSucceeded(value)
	}
}

type MoneyTransferCommand struct {
	CompositeBankAccountCommand
	from, to *BankAccount
	amount   int
}

func NewMoneyTransferCommand(from, to *BankAccount, amount int) *MoneyTransferCommand {
	c := &MoneyTransferCommand{from: from, to: to, amount: amount}
	c.commands = append(c.commands,
		NewBankAccountCommand(from, Withdraw, amount))
	c.commands = append(c.commands,
		NewBankAccountCommand(to, Deposit, amount))
	return c
}

func (m *MoneyTransferCommand) Call() {
	ok := true
	for _, cmd := range m.commands {
		if ok {
			cmd.Call()
			ok = cmd.Succeeded()
		} else {
			cmd.SetSucceeded(false)
		}
	}
}

func main() {
	from, to := BankAccount{100}, BankAccount{}

	cmd := NewMoneyTransferCommand(&from, &to, 100)

	fmt.Printf("from=%v to=%v\n", from, to)
	cmd.Call()
	fmt.Printf("from=%v to=%v\n", from, to)
	fmt.Println("Undoing...")
	cmd.Undo()
	fmt.Printf("from=%v to=%v\n", from, to)
}
