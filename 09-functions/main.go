package main

import (
	"fmt"
	"math"
)

func sayGreeting(n string) {
	fmt.Printf("Goodmorning, %v \n", n)
}

func sayGoodBye(n string) {
	fmt.Printf("Goodbye, %v \n", n)
}

func cycleNames(n []string, f func(string)) {
	for _, v := range n {
		f(v)
	}
}

func circleArea(r float64) float64 {
	return math.Pi * r * r
}

func main() {
	// sayGreeting("KC Samm")
	// sayGreeting("Michael")
	// sayGoodBye("Michael")

	// workers := []string{"Peter", "Nkem", "Victoria", "John"}

	// cycleNames(workers, sayGreeting)
	// cycleNames(workers, sayGoodBye)

	// Area of a Circle
	a1 := circleArea(7)
	a2 := circleArea(14.0)

	fmt.Printf("A1 = %0.1f \nA2 = %0.1f \n", a1, a2)
}