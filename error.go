package main

import (
	"errors"
	"fmt"
)

func Pembagian(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("Pembagian dengan nol")
	} else {
		return nilai / pembagi, nil
	}
}

func main() {
	result, error := Pembagian(100, 0)

	if error == nil {
		fmt.Printf("result = %v\n", result)
	}else {
		fmt.Printf("%v\n", error.Error())
	}
}
