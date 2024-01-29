package main

import "fmt"

func main() {
	name := "Fathar"

	if name == "Fariq" {
		fmt.Println("Hello Fariq")
	} else if name == "Fathar" {
		fmt.Println("Hello Fathar")
	} else {
		fmt.Println("Hai, nama kamu tidak ada dalam daftar")
	}

	// short statement
	if length := len(name); length > 5 {
		fmt.Println("Terlalu Panjang")
	} else {
		fmt.Println("Nama Sudah Benar")
	}
}
