package main

import (
	"fmt"

)

func main() {
	c := make(chan int)
	in := c
	out := make(chan int)

	go worker(in, out)

	in <- 0
	fmt.Println(<-out)
}

func worker(in, out  chan int) {
	v := <- in
	out <- 1 + v
}