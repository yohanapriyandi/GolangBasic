package main

import "fmt"

func main() {
	// for loops meggunakan statement
	for counter := 1; counter <= 10; counter++ {
		fmt.Println("perulangan ke", counter)
	}

	slice := []string{"yohan", "Nabilah", "fathar", "tsabita"}
	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	for i, v := range slice {
		fmt.Println("Index", i, "=", v)
	}

	person := make(map[string]string)
	person["name"] = "Yohan"
	person["title"] = "programmer"

	for k, v := range person {
		fmt.Println(k, "=", v)
	}

}
