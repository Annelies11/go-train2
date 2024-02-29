package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func changeCity(city *Address) {
	city.City = "Elmore"
}

func main() {
	gumball := Address{"", "Cartoon", "US"}

	changeCity(&gumball)

	fmt.Println(gumball)
}
