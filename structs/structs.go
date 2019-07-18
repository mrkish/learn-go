package main

import "fmt"

type Person struct {
	Name     string
	LastName string
	Age      int
}

func (p *Person) ChangeName(name string) string {
	p.Name = name
	return p.Name
}

func main() {

	var person1 Person
	person1.Name = "Joe"
	person1.LastName = "Bob"
	person1.Age = 21

	person2 := Person{
		Name:     "Bob",
		LastName: "Joe",
		Age:      12,
	}

	fmt.Println(person1)
	fmt.Println(person1.Name)
	fmt.Println(person2)
	person1.ChangeName("Hogsucker")
	fmt.Println(person1.Name)
}
