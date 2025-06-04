package main

import "fmt"

func main() {
	var number32 int32 = 32768
	var number64 int64 = int64(number32)
	var number16 int16 = int16(number64)

	fmt.Println(number32)
	fmt.Println(number64)
	fmt.Println(number16)


	var name = "Muhammad Ridha" 
	var m uint8 = name[0]
	var mString = string(m)

	fmt.Println(name)
	fmt.Println(m)
	fmt.Println(mString)
}
