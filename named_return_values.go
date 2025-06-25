package main

import "fmt"

func getCompleteName() (firstName, middleName, lastName string) {
	firstName = "Eko"
	middleName = "Kurniawan"
	lastName = "Khannedy"

	return firstName, middleName, lastName
}

func main() {
	fmt.Println(getCompleteName())
	a, b, c := getCompleteName()
	fmt.Println(a, b, c)
}
