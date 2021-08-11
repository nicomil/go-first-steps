package main

import (
	"fmt"
	"time"
)

func main() {
	// ex.1
	fmt.Println("Sum: ", sum(1, 2))
	// ex.2
	names := [...]string{"Nico", "Gracia", "Nahu"}
	names_slice := names[0:]
	printAll(names_slice)
	// ex.3
	numbers := [...]int{1, 34, 23, 22, 65, 34, 32}
	numbers_slice := numbers[0:]
	fmt.Println("Sum of arr numbers: ", sumArr(numbers_slice))
	// ex.4
	fmt.Println("Fact rec: ", factRec(6))
	fmt.Println("Fact proce: ", factProc(6))
}

func printAll(arr []string) {
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
}
func sum(x int, y int) int {
	return x + y
}

func sumArr(arr []int) int {
	sum := 0
	// similar to a foreach
	for _, elem := range arr {
		sum += elem
	}

	return sum
}

func replaceEven(arr []int) []int {
	// arr is a reference to original array
	for i := 0; i < len(arr); i++ {
		if arr[i]%2 == 0 {
			arr[i] = 0
		}
	}
	// range doesn't modify the actual value of the element because creates a copy called elem

	// for _,elem := range arr {
	// 	if elem%2 == 0 {
	// 		elem = 0
	// 	}
	// }
	return arr
}

func factRec(x int) int {
	if x == 1 {
		return x
	}
	return x * factRec(x-1)
}

func factProc(x int) int {
	var fact int = 1

	for i := 1; i <= x; i++ {
		fact *= i
	}
	return fact
}

// tests
func factsAvg(n int, f int) (time.Duration, time.Duration) {
	var avgArrRec []time.Duration
	var avgArrProc []time.Duration

	for i := 0; i < n; i++ {
		start := time.Now()
		factRec(f)
		duration := time.Since(start)
		avgArrRec = append(avgArrRec, duration)
		start2 := time.Now()
		factProc(f)
		duration2 := time.Since(start2)
		avgArrProc = append(avgArrProc, duration2)
	}
	avgRec := avg(avgArrRec)
	avgProc := avg(avgArrProc)
	return avgRec, avgProc
}

func avg(arr []time.Duration) time.Duration {
	var sum time.Duration = 0
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	return sum / time.Duration(len(arr))
}
