package main

import "fmt"

func hello(name string) string {
	return "Hello " + name
}

func main() {
	result := hello("Ridha")
	fmt.Println(result)
	fmt.Println(hello("eko"))
	fmt.Println(hello("brodi"))
}
