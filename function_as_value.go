package main

import "fmt"

func getGoodBye(name string) string {
	return "Good Bye, " + name
}

func main() {
	sayGoodBye := getGoodBye

	fmt.Println(sayGoodBye("Muhammad Ridha"))
}
