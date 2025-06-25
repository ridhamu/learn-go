package main

import "fmt"

type HasName interface {
	getName() string
}

func sayHelloAgain(hasName HasName) {
	fmt.Printf("Hello, %v!\n", hasName.getName())
}

type Person struct {
	Name string
}

func (person Person) getName() string {
	return person.Name
}

type Cat struct {
	Name string
}

func (cat Cat) getName() string {
	return cat.Name
}

func main() {
	person1 := Person{"Ridha"}
	sayHelloAgain(person1)

	cat1 := Cat{"Mimi"}
	sayHelloAgain(cat1)

}
