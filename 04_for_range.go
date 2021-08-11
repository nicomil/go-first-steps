package main

import ("fmt")

func main()  {
	// declaring i before
	var i int
	for i = 1; i<5; i++{
		fmt.Println(i)
	}
	// declaring j in for statement
	for j:=1; j<5; j++{
		fmt.Println(j)
	}

	// RANGES (similar to foreach)

	names := [3]string{"Nico", "Nahu", "Gracia"}
	// for index,vale := range array_name {
		//
	//}
	for _,name := range names {
		fmt.Printf("Member of the team: %s\n", name)
	}
}