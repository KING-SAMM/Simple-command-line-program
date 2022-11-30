/*
*** Working With Channels ***
(Goroutine to Goroutine Communication)
*/

package main

import (
	"fmt"
	"time"
)


func main() {
	c := make(chan string)
	go count("sheep", c)

	// Wrap in for loop to receive all messages
	/*
	for {
		msg, open := <- c // Receive from channel

		// Break out of loop if channel is no more open
		if !open {
			break
		}
		fmt.Println(msg)
	}
	*/

	// Alternative (cleaner) syntax
	for msg := range c {
		fmt.Println(msg)
	}
}

func count(animal string, c chan string) {
	for i:= 1; i <= 5; i++ {
		c <- animal // Send into channel
		time.Sleep(time.Millisecond * 500)
	}

	// Close channel when done sending to prevent deadlock error
	close(c)
}