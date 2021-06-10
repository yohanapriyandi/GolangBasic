package main

import "fmt"

func main() {
	var result = 10 + 10
	fmt.Println(result)

	var a = 10
	var b = 10
	var c = a * b
	fmt.Println(c)

	var x = 20
	var y = 10
	var z = x - y
	fmt.Println(z)

	var w = 100
	w += 50 // i = i + 50
	fmt.Println(w)

	var v = 100
	v -= 50 // i = i - 50
	fmt.Println(v)

	v++ // v = v + 1 => 51
	fmt.Println(v)

	v-- // v = v - 1 => 49
	fmt.Println(v)
}
