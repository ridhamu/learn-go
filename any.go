package main

func Ups() any {
	return 1
}

func EmptyInterfaceReturn() interface{} {
	return 0
}

func main() {
	Ups()
	EmptyInterfaceReturn()
}
