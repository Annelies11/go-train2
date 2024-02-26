package main

import "fmt"

type BL func(string) bool

func regUser(name string, blackList BL) {
	if blackList(name) {
		fmt.Println("Sorry,", name, ". You're Blocked!")
	} else {
		fmt.Println("Welcome, ", name)
	}
}

func main() {
	blacklist := func(name string) bool {
		return name == "Anjing"
	}
	regUser("Anjing", blacklist)

	regUser("Aris", func(name string) bool {
		return name == "Anjing"
	})
}
