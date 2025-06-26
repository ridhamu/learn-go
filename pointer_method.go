package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}

func main() {
	ridha := Man{"Muhammad Ridha"}
	ridha.Married()

	fmt.Println(ridha)
}
