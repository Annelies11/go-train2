package main

import "fmt"

func main() {
	counter := 0

	inc := func() {
		fmt.Println("Increment")
		counter++
	}

	inc()
	inc()
	inc()

	fmt.Println(counter)
}
