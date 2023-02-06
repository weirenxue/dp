package main

type Person struct {
	Name     string
	Age      int
	EyeCount int
}

func NewPerson(name string, age int) *Person {
	if age < 16 {
		return nil
	}
	return &Person{Name: name, Age: age, EyeCount: 2}
}

func main() {
	NewPerson("John", 23)
}
