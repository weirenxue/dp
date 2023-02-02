package main

import "fmt"

type BankAccount struct {
	balance int
}

func Deposit(ba *BankAccount, amount int) {
	fmt.Println("Depositing", amount)
	ba.balance += amount
}

func Withdraw(ba *BankAccount, amount int) {
	if ba.balance >= amount {
		fmt.Println("Withdrawing", amount)
		ba.balance -= amount
	}
}

func main() {
	ba := &BankAccount{}

	var commands []func()
	commands = append(commands, func() {
		Deposit(ba, 100)
	})
	commands = append(commands, func() {
		Withdraw(ba, 25)
	})

	for _, cmd := range commands {
		cmd()
	}
	fmt.Println(*ba)
}
