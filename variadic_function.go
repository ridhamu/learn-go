package main

import "fmt"

func sumAll(numbers ...int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	} 
	return sum
}


func main() {
	fmt.Println(sumAll(10, 10, 10))


	slice := []int {10, 10, 10, 10, 10}
	fmt.Println(sumAll(slice...))
}





