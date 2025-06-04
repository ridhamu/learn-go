package main

import "fmt"

func main() {
	var a = 10
	var b = 20
	var d = 5
	var e = 2
	var c = a + b - d * e

	fmt.Println(c)
	
	var i = 10
	i += 1
	fmt.Println("i = ", i)
	i -= 2
	fmt.Println("i = ", i)


	var j = 1
	fmt.Println("j = ", j)
	j++
	fmt.Println("j = ", j)
	j++
	fmt.Println("j = ", j)


	j--
	fmt.Println("j = ", j)
	j--
	fmt.Println("j = ", j)
	
}
