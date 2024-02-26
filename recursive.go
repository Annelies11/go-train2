package main

import "fmt"

func factorial(value int) int {
	res := 0

	for i := value; i > 0; i-- {
		res += i
	}
	return res
}

func realFact(val int) int {
	if val == 1 {
		return 1
	} else {
		return val + realFact(val-1)
	}
}

func main() {
	fmt.Println(factorial(4))
	fmt.Println(realFact(4))
}
