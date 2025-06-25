package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func (customer Customer) sayHelloTo(name string) {
	fmt.Printf("Hello %v👋, my name is %v!\n", name, customer.Name)
}

func main() {
	var person1 Customer
	person1.Name = "Ridha"
	person1.Address = "Gajayana"
	person1.Age = 22

	person1.sayHelloTo("Alejandro")

	fmt.Printf("%v\n", person1)
	fmt.Printf("%v\n", person1.Name)
	fmt.Printf("%v\n", person1.Address)
	fmt.Printf("%v\n", person1.Age)

	person2 := Customer{
		Name:    "Andi",
		Address: "Indonesia",
		Age:     25,
	}

	fmt.Printf("%v\n", person2)
	fmt.Printf("%v\n", person2.Name)
	fmt.Printf("%v\n", person2.Address)
	fmt.Printf("%v\n", person2.Age)

	person3 := Customer{"Alejandro", "Mexico", 31}
	fmt.Printf("%v\n", person3)
	fmt.Printf("%v\n", person3.Name)
	fmt.Printf("%v\n", person3.Address)
	fmt.Printf("%v\n", person3.Age)
}
