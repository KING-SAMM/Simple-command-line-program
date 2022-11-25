package main

import "fmt"

func main() {
	// Booleans
	y := 20
	x := 45 
	fmt.Println(x > y)
	fmt.Println(x < y)
	fmt.Println(y <= 20)
	fmt.Println(x >= 40)

	// Conditionals
	// x := 35
	// if x < 30 {
	// 	fmt.Println("Less than 30")
	// } else if x < 40 {
	// 	fmt.Println("Less than 40")
	// } else {
	// 	fmt.Println("Greater than 40")
	// }

	// Conditionals in loops
	// fruits := []string{"mango", "paw paw", "orange", "apple", "banana"}

	// for index, value := range fruits {
	// 	if index == 1 {
	// 		fmt.Println("Continuing at index", index)
	// 		continue
	// 	}
	// 	if index > 2 {
	// 		fmt.Println("Breaking at index", index)
	// 		break
	// 	}

	// 	fmt.Printf("Value at pos %v is %v \n", index, value)
	// }
}