package area

import "fmt"

func circleArea(r float) float64 {
	return math.PI * r ** 2
}

func main() {
	a1 := circleArea(7)
	a2 := circleArea(14.0)

	fmt.Printf("A1 = %0.1f \n A2 = %0.1f \n", a1, a2)
}