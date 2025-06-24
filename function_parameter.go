package main

import "fmt"

func sayHelloTo(firstName string, lastName string) {
	fmt.Printf("Hello, %v %v!", firstName, lastName);
}
func main(){
	sayHelloTo("Muhammad", "Ridha")
}
