// Send and receive through a channel within the same goroutine using a buffered channel


package main

import (
	"fmt"
)

func main() {
	c := make(chan string, 2)
	c <- "Hello"
	c <- "World"
	// c <- "Extra" // Blocking occurs here

	msg := <- c
	fmt.Println(msg)

	msg = <- c
	fmt.Println(msg)

	// msg = <- c
	// fmt.Println(msg) // Error for more items than buffer
}

// Error occurs only when tryng to send/receive n + 1 items on a buffer of n.