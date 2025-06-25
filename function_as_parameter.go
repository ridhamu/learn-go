package main

import "fmt"

type Filter func(string) string

func sayHelloWithFilter(name string, filter Filter) {
	filtered := filter(name)
	fmt.Printf("Hello, %v !\n", filtered)
}

func nameFilter(name string) string {
	if name == "Anjing" {
		return "***"
	} else if name == "Babi" {
		return "***"
	} else {
		return name
	}
}

func main() {
	sayHelloWithFilter("Ridha", nameFilter)

	filter := nameFilter
	sayHelloWithFilter("Anjing", filter)
	sayHelloWithFilter("Babi", filter)
}
