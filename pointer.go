package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	address1 := Address{"Lhokseumawe", "Aceh", "Indonesia"}
	address2 := address1 // by default is COPY


	address2.City = "Banda Aceh" // changing the copied value, not the address1.City original value
	fmt.Println(address1)
	fmt.Println(address2)

	fmt.Println("============pointer===========")
	var address3 Address = Address{"Malang", "Jawa Timur", "Indonesia"}
	var address4 *Address = &address3 // passing by reference or point to the same value in the memory

	address4.City = "Surabaya"
	fmt.Println(address3)
	fmt.Println(address3)
}
