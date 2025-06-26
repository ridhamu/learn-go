package main

import (
	"belajar-go-dasar/helper"
	"fmt"
)

func main() {
	name := helper.SayHello("Ridha")
	fmt.Println(name)

	fmt.Println(helper.Application)
	// fmt.Println(helper.version) // cannot access version because it only package access
	// fmt.Println(helper.sayGoodBye("Ridha")) // cannot access this either
}
