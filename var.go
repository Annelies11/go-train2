package main

import "fmt"

func main() {
	var name string
	var age int8

	name = "Aris Mahmudi"
	age = 25
	var wife = "I don't know"
	why := "Aris still want to focus to change himself"
	married := 30

	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(wife)
	fmt.Println(why)
	fmt.Println(married)

	var (
		book1 = "Bumi Manusia"
		book2 = "Twilight"
		book3 = "Laskar Pelangi"
	)

	const dream = "Being a lecturer"

	fmt.Println(book1)
	fmt.Println(book2)
	fmt.Println(book3)
	fmt.Println(dream)

}
