package main

import (
	"fmt"
	"time"
)

// step 1 without channel, main ends before gorouting could start
// func display() {
// 	time.Sleep(5 * time.Second)
// 	fmt.Println("Inside display()")
// }

// func main() {
// 	go display()
// 	fmt.Println("Inside main()")
// }

// step 2
func display(ch chan int) {
	time.Sleep(3 * time.Second)
	fmt.Println("Inside display()")
	ch <- 1234
}

func main() {
	ch := make(chan int)
	go display(ch)
	x := <-ch
	fmt.Println("Inside main()")
	fmt.Println("Printing x in main() after taking from channel:", x)

}
