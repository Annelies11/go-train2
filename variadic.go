package main

import "fmt"

func sumAll(numbers ...int) int {
	total := 0

	for _, n := range numbers {
		total += n
	}

	return total
}

func main() {
	res := sumAll(10, 10, 10)
	nn := []int{10, 10, 10, 10}

	fmt.Println(res)
	fmt.Println(sumAll(10, 10, 10, 10, 10))
	fmt.Println(nn)
}
