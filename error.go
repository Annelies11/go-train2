package main

import (
	"errors"
	"fmt"
)

func Div(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("Pembagian dengan 0")
	} else {
		return nilai / pembagi, nil
	}

}

func main() {
	hasil, err := Div(100, 0)
	if err == nil {
		fmt.Println("Hasil : ", hasil)
	} else {
		fmt.Println("Error : ", err.Error())
	}
}
