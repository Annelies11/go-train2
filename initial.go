package main

import (
	"03_third/database"
	_ "03_third/internal"
	"fmt"
)

func main() {
	fmt.Println(database.GetDatabase())
}
