package main

import "fmt"

func main() {
	var name1 = "Yohan"
	var name2 = "Yohan"

	var result bool = name1 == name2
	fmt.Println(result)

	var value1 = 1000
	var value2 = 2000

	fmt.Println(value1 > value2)
	fmt.Println(value1 < value2)
	fmt.Println(value1 == value2)
	fmt.Println(value1 != value2)

}
