package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}

func main() {
	aku := Man{"Aris"}
	aku.Married()

	fmt.Println(aku.Name)
}
