package main

import (
	"fmt"

)

const N = 100000

func main() {
	c := make(chan int)
	in := c
	var out chan int

	for i := 0; i < N; i++ {
		out = make(chan int)
		go worker(in, out)
		in = out
	}
	
	c <- 0
	fmt.Println(<-out)
}

func worker(in, out  chan int) {
	v := <- in
	out <- 1 + v
}