package main

import "fmt"

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
		return &validationError{"validation error"}
	}

	if id != "ridha" {
		return &notFoundError{"notFoundError"}
	}

	// ok
	return nil
}

func main() {
	test := SaveData("ridha", nil)

	if test != nil {
		// error
		if validationError, ok := test.(*validationError); ok {
			fmt.Println("validation error:", validationError.Message)
		} else if notFoundError, ok := test.(*notFoundError); ok {
			fmt.Println("Not Found Error:", notFoundError.Message)
		} else {
			fmt.Println("Unknown error:", test.Error())
		}
	} else {
		// success
		fmt.Println("success!")
	}
}
