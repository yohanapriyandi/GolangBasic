package main

import (
	"fmt"
)

func main() {
	person := map[string]string{
		"name":    "Yohan",
		"address": "Kuningan",
	}
	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])

	var book map[string]string = make(map[string]string)
	book["title"] = "Belajar Golang"
	book["author"] = "yohan"
	book["ups"] = "what"

	// delete(book, "ups")
	fmt.Println(book)
	fmt.Println(len(book))
}
