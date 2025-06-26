
package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	fmt.Println("============pointer===========")
	var address1 Address = Address{"Malang", "Jawa Timur", "Indonesia"}
	var address2 *Address = &address1 // by reference

	// address2 = &Address{"Jakarta", "Jakarta", "Indonesia"}
	*address2 = Address{"Jakarta", "Jakarta", "Indonesia"}

	fmt.Println(address1)
	fmt.Println(address2)
}
