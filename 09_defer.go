package main

import "fmt"

// defer keyword wait the containing function to complete the execution and then executes
// in case of stacking defer calls they will execute LIFO
// stack
// 3 -> first
// 2 -> second
// 1 -> third
func main() {
	defer display(1)
	defer display(2)
	defer display(3)
	fmt.Println(4)
	// output
	// 4
	// 3
	// 2
	// 1
}

func display(num int) {
	fmt.Println(num)
}
