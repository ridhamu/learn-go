package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	var address1 *Address = &Address{}
	var address2 *Address = address1

	address2.Country = "Singapore"

	fmt.Println(address1)
	fmt.Println(address2)

	address3 := new(Address)
	address4 := address3

	address4.Country = "Indonesia"

	fmt.Println(address3)
	fmt.Println(address4)
}
