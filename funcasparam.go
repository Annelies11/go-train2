package main

import "fmt"

type Filter func(string) string

func sayHello(name string, filter Filter) {
	filteredName := filter(name)
	fmt.Println("Hello, ", filteredName)
}

func spamFilter(name string) string {
	if name == "Anjing" {
		return "..."
	} else {
		return name
	}
}

func main() {
	sayHello("Aris", spamFilter)

	filter := spamFilter
	sayHello("Anjing", filter)

	sayHello("Anjing", spamFilter)
}
