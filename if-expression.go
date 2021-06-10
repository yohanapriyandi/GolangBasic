package main

import "fmt"

func main() {
	name := "thar"

	if name == "Eko" {
		fmt.Println("Hello Eko")
	} else if name == "Fathar" {
		fmt.Println("Hello Fathar")
	} else {
		fmt.Println("Hai")
	}

	if length := len(name); length > 5 { // short statement
		fmt.Println("Terlalu Panjang")
	} else {
		fmt.Println("Nama Sudah Benar")
	}
}
