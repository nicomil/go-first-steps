package main

import "fmt"

// defining the structure
type user struct {
	name     string
	lastname string
	age      int
}

func main() {

	// creating a variable of type user
	var nico user
	nico.name = "Nico"
	nico.lastname = "Mil"
	nico.age = 33

	gracia := user{"Gracia", "Ruiz", 35}

	// displayUser(nico)
	// displayUser(gracia)

	nico.display()
	gracia.display()
}

// normal function
func displayUser(u user) {
	fmt.Println("Name: ", u.name)
	fmt.Println("Lastname: ", u.lastname)
	fmt.Println("Age: ", u.age)
	fmt.Printf("\n")
}

// method
func (u user) display() {
	displayUser(u)
}
