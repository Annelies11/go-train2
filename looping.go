package main

import "fmt"

func main() {

	for i := 0; i < 20; i++ {

		if i < 19 {
			fmt.Print(i+1, " - ")
		} else {
			fmt.Print(i + 1)
		}
	}

	names := []string{"Aris", "Mahmudi", "Mavis"}
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	for i, name := range names {
		fmt.Println("Index : ", i, "Name : ", name)
	}

	for _, name := range names {
		fmt.Println(name)
	}
}
