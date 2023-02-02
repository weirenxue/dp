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

func NewBankAccountCommand(account *BankAccount, action Action, amount int) *BankAccountCommand {
	return &BankAccountCommand{account: account, action: action, amount: amount}
}

func main() {
	ba := BankAccount{}

	cmd := NewBankAccountCommand(&ba, Deposit, 100)
	cmd.Call()

	cmd2 := NewBankAccountCommand(&ba, Withdraw, 150)
	cmd2.Call()

	cmd3 := NewBankAccountCommand(&ba, Withdraw, 1000)
	cmd3.Call()

	cmd3.Undo()
	cmd2.Undo()
	cmd.Undo()

	fmt.Println(ba)
}
