package main

import "fmt"


func main() {
	name := "Muhammad"

	switch name {
	case "Ridha": 
		fmt.Printf("Halo, Ridha!\n")
	case "Muhammad": 
		fmt.Printf("Halo, Muhammad!\n")
	default: 
		fmt.Printf("Halo, Boleh Kenalan ?\n")
	}

	// short statement in switch
	switch panjangNama := len(name); panjangNama > 5 {
	case true: 
	fmt.Printf("Panjang nama lebih dari 5 karakter\n")
	case false:
	fmt.Printf("Panjang nama kurang dari 5 karakter\n")
	}


	// empty switch statement
	name = "dha"
	length := len(name)
	switch {
	case length > 10:
		fmt.Printf("nama terlalu panjang")
	case length > 5:
		fmt.Printf("nama lumayan panjang")
	default:
		fmt.Printf("nama sudah pas")
	}
}
