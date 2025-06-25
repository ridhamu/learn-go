package main

import "fmt"

func endApp() {
	fmt.Printf("Ending application...\n")
	if message := recover(); message != nil {
		fmt.Printf("Panic Message: %v\n", message)
	}
}

func runApp(error bool) {
	defer endApp()

	if error {
		panic("SERVER ERROR")
	}

	fmt.Printf("Server not Error :) \n")
}

func main() {
	runApp(true)
}
