package main

import "fmt"

func main() {
	
	// While loop equivalent
	// x := 0

	// for x < 5 {
	// 	fmt.Println("The value of x is:", x)
	// 	x++
	// }


	// for loop equivalent
	// for i := 0; i < 5; i++ {
	// 	fmt.Println("The value of i is:", i)
	// }

	fruits := []string{"mango", "apple", "paw paw", "orange"}

	// for i := 0; i < len(fruits); i++ {
	// 	fmt.Printf("Fruit %v is %v \n", i, fruits[i])
	// }

	// for-in loop equivalent
	// for index, value := range fruits {
	// 	fmt.Printf("Fruit %v is %v \n", index, value)
	// }

	// Omtting a variable with underscore
	// for _, value := range fruits {
	// 	fmt.Printf("Eating %v \n", value)
	// }

	// Using only index
	for index := range fruits {
		fmt.Printf("Fruit %v is %v \n", index, fruits[index])
	}

	
}