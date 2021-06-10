package main

import "fmt"

func main() {
	var name [3]string

	name[0] = "Yohan"
	name[1] = "Fithriya"
	name[2] = "Fathar"

	fmt.Println(name[0])
	fmt.Println(name[1])
	fmt.Println(name[2])

	var values = [3]int{
		10, 20, 30,
	}
	fmt.Println(values)

	fmt.Println(len(name))
	fmt.Println(len(values))

}
