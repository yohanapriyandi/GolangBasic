package main

import "fmt"

func main() {
	name := "Fathar Dhabit Adz-Dzaky"

	switch name {
	case "Yohan":
		fmt.Println("Hai Yohan")
	case "Fathar":
		fmt.Println("Hai Fathar")
	default:
		fmt.Println("Hai")
	}

	// short statement
	// switch length := len(name); length > 5 {
	// case true:
	// 	fmt.Println("Terlalu Panjang")
	// case false:
	// 	fmt.Println("Nama sudah benar")
	// }

	length := len(name)

	switch {
	case length > 10:
		fmt.Println("Nama terlalu Panjang")
	case length > 5:
		fmt.Println("Nama lumayan Panjang")
	default:
		fmt.Println("Nama sudah benar")
	}
}
