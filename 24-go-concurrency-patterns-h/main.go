// The SELECT Pattern: The Problem

package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			c1 <- "Every 500 milliseconds"
			time.Sleep(time.Millisecond * 500)
		}
	}()
	go func() {
		for {
			c2 <- "Every 2 seconds"
			time.Sleep(time.Second * 2)
		}
	}()

	for {
		fmt.Println(<-c1) // Is blocked by c2 and output every 2 s as a result
		fmt.Println(<-c2) //Blocks c1 from which we can receive quicker
	}
}