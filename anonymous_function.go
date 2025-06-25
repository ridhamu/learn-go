package main

import "fmt"

type BlackList func(string) bool

func registerUser(name string, blackList BlackList) {
	if blackList(name) {
		fmt.Printf("You are blocked, %v!\n", name)
	} else {
		fmt.Printf("Welcome, %v!\n", name)
	}
}

func main() {
	blackList := func(name string) bool {
		switch name {
		case "Andi":
			return true
		case "Budi":
			return true
		case "Charles":
			return true
		default:
			return false
		}
	}

	registerUser("Ridha", blackList)
	registerUser("Andi", blackList)
	registerUser("Budi", blackList)

	// or just define the function inside the func param
	registerUser("Muhammad", func(name string) bool {
		switch name {
		case "Andi":
			return true
		case "Budi":
			return true
		case "Charles":
			return true
		default:
			return false
		}
	})
}
