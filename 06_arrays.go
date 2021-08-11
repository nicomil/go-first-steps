package main

import (
	"fmt"
)

// Arrays
// var arrayname [size] type (declaring)
// arrayname [index] = value (assigning to element)
// arrayname := [size] type {value_0,value_1,…,value_size-1}
// arrayname :=  […] type {value_0,value_1,…,value_size-1}

// len(arrayname) to get le length
func main() {
	var threePigs [3]string

	threePigs[0] = "Pig 1"
	threePigs[1] = "Pig 2"
	threePigs[2] = "Pig 3"

	fmt.Println(threePigs)
	fmt.Println(len(threePigs))
	fmt.Println(threePigs[2])

	something := [...]int{2, 3, 34, 23, 4, 3, 34}

	fmt.Println(something)
	fmt.Println(len(something))
	fmt.Println(fmt.Sprintf("%#v", something))
	// printArr(something)
}

func printArr(arr []int) {
	for _, elem := range arr {
		fmt.Println(elem)
	}
}
