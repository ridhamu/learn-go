package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func changeCountryToIndonesia(address *Address) {
	address.Country = "Indonesia"
}

func main() {
	address1 := Address{}

	changeCountryToIndonesia(&address1)
	fmt.Println(address1)
}
