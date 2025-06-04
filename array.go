package main

import "fmt"

func main() {
	var fullname [2]string
	
	fullname[0] = "Muhammad"
	fullname[1] = "Ridha"

	fmt.Println(fullname[0])
	fmt.Println(fullname[1])

	var values = [3]int {
		1,
		2, 
		3,
	}

	fmt.Println("array [3]int adalah", values)


	// we also able to initiate an array without its size at first, but must insert the value on it
	var value2 = [...]int {1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}

	fmt.Printf("value2 [...]int with size of %v = %v \n", len(value2), value2)

}
