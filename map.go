package main

import "fmt"

func main() {
	person := map[string]string{
		"name": "Muhammad Ridha",
		"addres": "Jl. Gajayana gg.6 No.577C",
	}
	fmt.Printf("person = %v\n", person)
	fmt.Printf("person['name'] = %v\n", person["name"])
	fmt.Printf("person['addres'] = %v\n", person["addres"])


	book := make(map[string]string)
	book["title"] = "Harry Pottter"
	book["author"] = "J.K Rowling"

	fmt.Printf("book = %v\n", book)
	fmt.Printf("book['title'] = %v\n", book["title"])
	fmt.Printf("book['author'] = %v\n", book["author"])
//	fmt.Printf("book['false'] = %v\n", book["false"]) // this is error, because we tried to acces false key inside the map
}






