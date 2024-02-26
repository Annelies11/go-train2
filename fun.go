package main

import "fmt"

func say(name string) string {
	return ("Hello, " + name)
}

func sayy() (string, string) {
	return "UJUJU", "KUKUKU"
}

func getCom() (first, middle, last string) {
	first = "Aris"
	middle = "Mah"
	last = "Mudi"

	return first, middle, last
}

func main() {
	res := say("Aris")
	fmt.Println(res)
	ros, rem := sayy()
	fmt.Println(ros)
	fmt.Println(rem)

	fi, mid, la := getCom()
	fmt.Println(fi, mid, la)
}
