package main

import "fmt"

func factorialManual(number int) int {
	result := 1

	for i := number; i > 0; i-- {
		result *= i
	}
	return result
}

func factorial(number int) int {
	if number <= 1 {
		return 1
	} else {
		return number * factorial(number-1)
	}
}

func main() {
	result := 10 * 9 * 8 * 7 * 6 * 5 * 4 * 3 * 2 * 1
	fmt.Printf("correct result = %v\n", result)
	fmt.Printf("%v\n", factorial(10))
	fmt.Printf("%v\n", factorialManual(-10))
}
