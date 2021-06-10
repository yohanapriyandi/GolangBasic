package main

import (
	"fmt"
	"reflect"
)

func main() {
	// fmt.Println("Hello world")
	num1, num2 := 10, true
	// num = 12
	fmt.Println(reflect.TypeOf(num1))
	fmt.Println(reflect.TypeOf(num2))
}
