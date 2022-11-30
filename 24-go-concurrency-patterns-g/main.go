// Send and receive through channels within the same goroutine using buffered channels 

package main

import (
	"fmt"
)

func main() {
	c := make(chan string, 2)
	c <- "Hello"

	msg := <- c
	fmt.Println(msg)
}