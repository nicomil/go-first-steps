package main

import (
	"fmt"
)

// SLICE
// var slice_name [] type = array_name[start:end]
// slice_name is a reference to array_name
// values changed in slice_name change in array_name
func main() {
	// declaring array with undefined lenght
	a := [...]string{"one", "two", "three", "four", "five"}
	fmt.Println("Array after creation:", a)

	b := a[1:4] //created a slice named b
	fmt.Println("Slice after creation:", b)

	b[0] = "changed" // changed the slice data
	fmt.Println("Slice after modifying:", b)
	fmt.Println("Array after slice modification:", a)

	teamA := [...]string{"James", "Sam", "Nico"}
	fmt.Println("TeamA:", teamA)
	// slice_a := teamA[0:len(teamA)]
	slice_a := teamA[:]

	teamB := [...]string{"Alex", "Diana", "Leila"}
	fmt.Println("TeamB:", teamB)
	slice_b := teamB[0:len(teamB)]

	slice_a = append(slice_a, slice_b...)
	all := slice_a
	fmt.Println(all)
	all[0] = "Rino"
	fmt.Println(slice_a)
}
