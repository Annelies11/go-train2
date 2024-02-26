package main

import "fmt"

func main() {
	me := map[string]string{
		"name": "Aris",
		"age":  "Twenty Five",
	}
	me["address"] = "Tulungagung"

	fmt.Println(me["name"])
	fmt.Println(me["age"])
	fmt.Println(me["address"])
	me["age"] = ""

	fmt.Println(me)

	
}
