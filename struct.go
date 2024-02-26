package main

import "fmt"

type HasName interface {
	getName() string
}

func hello(value HasName) {
	fmt.Println("Hello,", value.getName())
}

type Person struct {
	Name string
}

func (person Person) getName() string {
	return person.Name
}

type Customer struct {
	Name, Address string
	Age           int
}

func (custom Customer) seyHello(name string) {
	fmt.Println("Hello, ", name, " My name is ", custom.Name)
}

func main() {
	budi := Customer{"Budi", "Mars", 40}

	budi.seyHello("Simple")

	bego := Person{Name: "Boogie"}

	hello(bego)
}
