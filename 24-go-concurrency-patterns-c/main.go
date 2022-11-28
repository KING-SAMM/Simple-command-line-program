package main
import (
	"fmt"
	"time"
)

// func main() {
// 	count("sheep")
// 	count("fish")
// }


// func main() {
// 	go count("sheep")
// 	count("fish")
// }


// func main() {
// 	go count("sheep")
// 	go count("fish")
// }


func main() {
	go count("sheep")
	go count("fish")

	// time.Sleep(time.Second * 2)
	// fmt.Scanln()
}

func main() {
	go count("sheep")
	go count("fish")

	fmt.Scanln()
}

func count(animal string) {
	for i := 0; true; i++ {
		fmt.Println(i, animal)
		time.Sleep(time.Millisecond * 500)
	}
}