package main

import "fmt"

func main() {
	var name string = "ridha"

	if name == "ridha" {
		fmt.Printf("halo, %v👋!\n", name)
	}else if name == "budi" {
		fmt.Printf("halo, %v👋!\n", name)
	}else {
		fmt.Printf("halo, boleh kenalan?\n")
	}


	//if short statement
	if length := len(name); length > 5 {
		fmt.Printf("nama lebih dari 5 kata\n")
	}else {
		fmt.Printf("nama kurang dari 5 kata\n")
	}
}
