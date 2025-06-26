package main

import "fmt"

func random() any {
	return "OK"
}

func main() {
	var test any = random()
	// var resultString string = test.(string)

	// fmt.Printf("test = %v\ttype:%T\n", test, test)
	// fmt.Printf("resultString = %v\ttype:%T\n", resultString, resultString)

	// if we tried to assert type from any to int that actually a int it would cause a panic
	// var resultInt int = test.(int)
	// fmt.Printf("resultInt = %v\n", resultInt)

	// so what we do is we use switch to safely assert possible type
	// first, we store the type inside a variable
	// then we switch case from there
	switch value := test.(type) {
	case string:
		fmt.Printf("string =\t%v\n", value)
	case int:
		fmt.Printf("int =\t%v\n", value)
	default:
		fmt.Printf("unknown =\t%v\n", value)
	}
}
