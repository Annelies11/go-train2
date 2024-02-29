package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	var address1 Address = Address{"Kediri", "Jatim", "Indonesia"}
	var address2 *Address = &address1

	fmt.Println(address1)
	fmt.Println(address2)

	address2.City = "Tulungagung"

	*address2 = Address{"Pyongyang", "Ambon", "Uganda"}
	fmt.Println(address1)
	fmt.Println(address2)

	add1 := new(Address)
	add2 := add1

	add2.City = "Saranjana"
	fmt.Println(add1)
	fmt.Println(add2)
}
