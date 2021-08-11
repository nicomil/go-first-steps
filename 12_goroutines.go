package main

import (
	"fmt"
	"time"
)

func display() {
	for i := 0; i < 5; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println("In display")
	}
}

func main() {
	//invoking the goroutine display()
	go display()
	//The main() continues without waiting for display()
	for i := 0; i < 5; i++ {
		// without timer main execution ends before display can start, with timer they run together
		time.Sleep(3 * time.Second)
		fmt.Println("In main")
	}
	// In display
	// In display
	// In main
	// In display
	// In display
	// In display
	// In main
	// In main
	// In main
	// In main
}
