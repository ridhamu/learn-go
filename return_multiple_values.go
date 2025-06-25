package main

import "fmt"

func getFullName() (string, string) {
	return "Muhammad", "Ridha"
}

func main() {
	firstName, _ := getFullName()
	fmt.Println(firstName)
}
