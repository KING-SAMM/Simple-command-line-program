package main

import "fmt"

func main() {
	
	// ARRAYS
	var ages [3]int = [3]int{12, 14, 23}
	fruits := [4]string{"apple", "orange", "paw paw", "pear"}

	fmt.Println(ages, len(ages))
	fmt.Println(fruits, len(fruits))


	// SLICES
	var scores = []int{67, 83, 48}
	ingredients := []string{"pepper", "egg", "salt", "sugar"}

	scores = append(scores, 90)

	ingredients[1] = "green pea"
	ingredients = append(ingredients, "water")

	fmt.Println(scores, len(scores))
	fmt.Println(ingredients, len(ingredients))

	// RANGE
	rangeOne := ingredients[2:4]
	rangeTwo := ingredients[2:]
	rangeThree := ingredients[:4]

	fmt.Println(rangeOne, rangeTwo, rangeThree)

}