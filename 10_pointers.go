package main

import "fmt"

func main() {
	a := 20
	// & operator to get the address of the variable
	fmt.Println("Address: ", &a) // Address:  0xc000118000
	fmt.Println("Value: ", a)    // Value:  20

	// a pointer variable store the address of another variable
	var aPointer *int = &a

	fmt.Println(aPointer)  // 0xc000118000
	fmt.Println(*aPointer) // 20

	*aPointer++

	fmt.Println(a)

	num := 30

	var numP *int = &num

	fmt.Println(*numP)
	num++
	fmt.Println(*numP)
	*numP++
	fmt.Println(*numP)

	// in case of array just use slice instead of pointer, isn't a good way and code becomes more complicated to read
	arr := []int{78, 89, 45, 56, 14}
	updateArrSlice(arr[:]) // best option
	fmt.Println(arr)
}

func updateArrPointer(arr *[5]int) {
	(*arr)[0] = 1000
}
func updateArrPointer2(arr *[5]int) {
	arr[0] = 1000
}
func updateArrSlice(arr []int) {
	arr[0] = 1000
}
