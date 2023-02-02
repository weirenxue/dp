package main

import "fmt"

type Memento struct {
	Balance int
}

type BankAccount struct {
	balance int
	changes []*Memento
	current int
}

func NewBankAccount(balance int) *BankAccount {
	ba := &BankAccount{balance: balance}
	m := &Memento{Balance: balance}
	ba.changes = append(ba.changes, m)
	return ba
}

func (b *BankAccount) String() string {
	return fmt.Sprint("Balance = $", b.balance, ", current = ", b.current)
}

func (b *BankAccount) Deposit(amount int) *Memento {
	b.balance += amount
	m := &Memento{Balance: b.balance}

	if b.current != len(b.changes)-1 {
		b.changes = b.changes[:b.current+1]
	}

	b.changes = append(b.changes, m)
	b.current++

	fmt.Println("Deposited", amount, "\b, balance is now", b.balance)

	return m
}

func (b *BankAccount) Undo() *Memento {
	if b.current > 0 {
		b.current--
		m := b.changes[b.current]
		b.balance = m.Balance
		return m
	}
	return nil
}

func (b *BankAccount) Redo() *Memento {
	if b.current+1 < len(b.changes) {
		b.current++
		m := b.changes[b.current]
		b.balance = m.Balance
		return m
	}
	return nil
}

func main() {
	ba := NewBankAccount(100)
	fmt.Println(ba)

	ba.Deposit(50)
	ba.Deposit(25)
	fmt.Println(ba)

	ba.Undo()
	fmt.Println("Undo 1:", ba)
	ba.Undo()
	fmt.Println("Undo 2:", ba)
	ba.Undo()
	fmt.Println("Undo 3:", ba)

	ba.Deposit(1000)
	ba.Deposit(2000)
	fmt.Println(ba)

	ba.Redo()
	fmt.Println("Redo 1:", ba)
	ba.Undo()
	fmt.Println("Undo 4:", ba)
	ba.Redo()
	fmt.Println("Redo 2:", ba)
}

// type Memento struct {
// 	Balance int
// }

// type BackAccount struct {
// 	balance int
// 	changes []*Memento
// 	current int
// }

// func (b *BackAccount) String() string {
// 	return fmt.Sprint("Balance = $", b.balance, ", current = ", b.current)
// }

// func NewBankAccount(balance int) *BackAccount {
// 	b := &BackAccount{balance: balance}
// 	b.changes = append(b.changes, &Memento{Balance: balance})
// 	return b
// }

// func (b *BackAccount) Deposit(amount int) *Memento {
// 	b.balance += amount
// 	m := &Memento{Balance: b.balance}
// 	if b.current == len(b.changes)-1 {
// 		b.changes = append(b.changes, m)
// 		b.current++
// 	} else {
// 		b.current++
// 		b.changes = b.changes[:b.current]
// 		b.changes = append(b.changes, m)
// 	}
// 	fmt.Println("Deposited", amount,
// 		"\b, balance is now", b.balance)
// 	return m
// }

// func (b *BackAccount) Restore(m *Memento) {
// 	if m != nil {
// 		b.balance = m.Balance
// 		b.changes = append(b.changes, m)
// 		b.current = len(b.changes) - 1
// 	}
// }

// func (b *BackAccount) Undo() *Memento {
// 	if b.current > 0 {
// 		b.current--
// 		m := b.changes[b.current]
// 		b.balance = m.Balance
// 		return m
// 	}
// 	return nil
// }

// func (b *BackAccount) Redo() *Memento {
// 	if b.current+1 < len(b.changes) {
// 		b.current++
// 		m := b.changes[b.current]
// 		b.balance = m.Balance
// 		return m
// 	}
// 	return nil
// }
