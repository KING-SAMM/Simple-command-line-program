package main
import (
	"fmt"
	"time"
)

func main() {
	go count("sheep")
	count("fish")
}

func count(animal string) {
	for i := 0; true; i++ {
		fmt.Println(i, animal)
		time.Sleep(time.Millisecond * 500)
	}
}