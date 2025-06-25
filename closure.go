package main

import "fmt"

func main() {
	counter := 0
	fmt.Printf("counter = %v\n", counter)
	/*
	really really longggggg code
	*/

	increment := func() {
		fmt.Printf("incrementing...\n")
		counter++
	}

	increment()
	increment()

	fmt.Printf("counter = %v\n", counter)
}
