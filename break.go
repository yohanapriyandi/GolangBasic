package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		if i == 3 {
			break
		}
		fmt.Println("perulangan ke", i)
	}
}
