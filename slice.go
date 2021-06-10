package main

import "fmt"

func main() {
	var months = [...]string{
		"Januari",
		"Pebruari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}

	var slice1 = months[4:7]
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	newSlice := make([]string, 2, 5) //membuat slice baru

	newSlice[0] = "Arsenal"
	newSlice[1] = "Liverpool"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	copySlice := make([]string, len(newSlice), cap(newSlice)) // duplikay sice
	copy(copySlice, newSlice)
	fmt.Println(copySlice)

	bentukArray := [3]int{1, 2, 3}
	bentukSlice := []int{1, 2, 3}

	fmt.Println(bentukArray)
	fmt.Println(bentukSlice)

}
