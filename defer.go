package main

import "fmt" 

func logging(fnName string) {
	fmt.Printf("Success Calling function %v\n", fnName)
}

func runApplication() {
	defer logging("runApplication")

	fmt.Printf("Running Application...\n")
}

func main() {
	runApplication()
}
