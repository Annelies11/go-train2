package main

import "fmt"

func rand() any {
	return true
}

func main() {
	var result any = rand()
	// var resString string = result.(string)
	// fmt.Println(resString)

	// var resInt int = result.(int)
	// fmt.Println(resInt)
	switch value := result.(type) {
	case string:
		fmt.Println("String,", value)
	case int:
		fmt.Println("Int,", value)
	default:
		fmt.Println("Unknown")
	}
}
