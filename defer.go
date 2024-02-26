package main

import "fmt"

func logg() {
	fmt.Println("This is Loging")
}

func runApp() {
	defer logg()
	fmt.Println("This is runApp")
	fmt.Println("...")
	fmt.Println("After this block done")
}

func endApp() {
	fmt.Println("End App")
	rec := recover()
	fmt.Println("What panic?", rec)
}

func destroyer(error bool) {
	defer endApp()
	if error {
		panic("OOOPS")
	}

}

func main() {
	runApp()
	destroyer(true)
}
