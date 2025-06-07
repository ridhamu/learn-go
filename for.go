package main

import "fmt"

func main() {
	counter := 1

	// the default for without any init and post statement
	for counter <= 10 {
		fmt.Printf("perulangan ke-%v\n", counter)
		counter++
	}

	fmt.Printf("\n============next============\n")
	// for loop with init and post statement
	for i := 1; i <= 10; i++ {
		fmt.Printf("perulangan ke-%v\n", i)
	}


	fmt.Printf("\n============next============\n")
	// for range
	// iterate through an array
	array := [...]string {"banana", "watermelon", "apple"}

	for index, value := range array {
		fmt.Printf("index-%v = %v\n", index, value)
	}

	fmt.Printf("\n============next============\n")
	// iterate through a slice
	slice := []int {1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, value := range slice {
		fmt.Printf("nilainya adalah %v\n", value)
	}

	fmt.Printf("\n============next============\n")
	//iterate through a map
	person1 := map[string]string {
		"name": "Muhammad Ridha",
		"address": "Jl. Gajayana gg.6 no. 577C",
		"age": "22",
	}

	for key, value := range person1 {
		fmt.Printf("%v: %v\n", key, value)
	}
}











