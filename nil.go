package main

import "fmt"

func NewMap(name string) map[string]string {
	if name == ""{
		return nil
	}else {
		return map[string]string{
			"name": name,
		}
	}
}

func main() {
	data := NewMap("ridha")

	if data == nil {
		fmt.Println("data nil!")
	}else {
		fmt.Printf("%v\n", data)
	}
}
