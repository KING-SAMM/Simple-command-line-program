// Send and receive through a channel within the same goroutine will NOT work 

package main

import (
	"fmt"
)

func main() {
	c := make(chan string)
	c <- "Hello"

	msg := <- c
	fmt.Println(msg)
}