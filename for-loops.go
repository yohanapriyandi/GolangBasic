package main

import "fmt"

func main() {
	// for loops meggunakan statement
	fmt.Println("=================for loops using statement====================")
	for counter := 1; counter <= 10; counter++ {
		fmt.Println("perulangan ke", counter)
	}
	fmt.Println("=======================================")
	slice := []string{"yohan", "Nabilah", "fathar", "tsabita"}
	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}
	fmt.Println("=======================================")
	for i, v := range slice {
		fmt.Println("Index", i, "=", v)
	}
	fmt.Println("=======================================")
	person := make(map[string]string)
	person["name"] = "Yohan"
	person["title"] = "programmer"

	for k, v := range person {
		fmt.Println(k, "=", v)
	}

}
