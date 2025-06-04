package main

import "fmt"

func main() {
	// deklarasi alias dari suatu tipe data
	type NoKtp string

	var ktpRidha NoKtp = "117302250203"

	//in case need to convert data to typedeclration vice versa 
	var dataLama = "2222222222"
	var dataDiperbaharui = NoKtp(dataLama)


	fmt.Println(ktpRidha)
	fmt.Println(dataDiperbaharui)
}
