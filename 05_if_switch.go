package main

import ("fmt")

func main()  {
	// if else
	var num int = 13

	if num % 2 == 0{
		fmt.Println("Numero par")
	}else{
		fmt.Println("Numero impar")
	}

	// switch 
	var color string = "blue"

	switch color {
	case "green": 
				fmt.Println("You're the grinch")
	case "blue": 
				fmt.Println("You're the puffo")	
	case "red": 
				fmt.Println("You're the volcano")
	default: 
				fmt.Println("You're nothing")
	}
}