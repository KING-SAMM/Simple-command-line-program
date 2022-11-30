// The SELECT Pattern: The Solution
// Here, we receive as often as c1 is ready without blocking from c2

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
		select {
		case msg1 := <- c1:
			fmt.Println(msg1) 
		case msg2 := <- c2:
			fmt.Println(msg2)
		}
	}
}