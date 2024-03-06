package main

import (
	"fmt"
)

type validationError struct {
	Message string
}

func (v *validationError) Error() string {
	return v.Message
}

type notFoundError struct {
	Message string
}

func (n *notFoundError) Error() string {
	return n.Message
}

func SaveData(id string, data any) error {
	if id == "" {
		return &validationError{Message: "Validation Error"}
	}
	if id != "Aris" {
		return &notFoundError{Message: "Data Not Found"}
	}
	return nil
}

func main() {
	err := SaveData("", 12)
	if err != nil {
		// if validationError, ok := err.(*validationError); ok {
		// 	fmt.Println("Validation error : ", validationError.Error())
		// } else if notFoundError, ok := err.(*notFoundError); ok {
		// 	fmt.Println("Not Found Error : ", notFoundError.Error())
		// } else {
		// 	fmt.Println("Unknown Error : ", err.Error())
		// }
		switch finalError := err.(type) {
		case *validationError:
			fmt.Println("Validation error : ", finalError.Error())
		case *notFoundError:
			fmt.Println("Validation error : ", finalError.Error())
		default:
			fmt.Println("Unknown Error : ", finalError.Error())
		}
	} else {
		fmt.Println("Sukses!")
	}
}
