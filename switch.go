package main

import "fmt"

func main() {
	num := 6

	switch num {
	case 1:
		fmt.Println("A")
	case 2:
		fmt.Println("B")
	case 3:
		fmt.Println("C")
	case 4:
		fmt.Println("D")
	default:
		fmt.Println("ZZZ")
	}
}
