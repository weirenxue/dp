package main

import "fmt"

type Employee struct {
	Name, Position string
	AnnualIncome   int
}

func NewEmployeeFactory(position string, annualIncome int) func(name string) *Employee {
	return func(name string) *Employee {
		return &Employee{Name: name, Position: position, AnnualIncome: annualIncome}
	}
}

type EmployeeFactory struct {
	Position     string
	AnnualIncome int
}

func NewEmployeeFactory2(position string, annualIncome int) *EmployeeFactory {
	return &EmployeeFactory{Position: position, AnnualIncome: annualIncome}
}

func (e *EmployeeFactory) Create(name string) *Employee {
	return &Employee{Name: name, Position: e.Position, AnnualIncome: e.AnnualIncome}
}

func main() {
	developerFactory := NewEmployeeFactory("developer", 60000)
	managerFactory := NewEmployeeFactory("manager", 80000)
	john := developerFactory("John")
	matt := managerFactory("Matt")
	fmt.Println(*john)
	fmt.Println(*matt)

	bossFactory := NewEmployeeFactory2("CEO", 100000)
	bossFactory.AnnualIncome = 110000
	merry := bossFactory.Create("Merry")
	fmt.Println(*merry)
}
