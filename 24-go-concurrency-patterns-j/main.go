/* 
* WORKER POOLS PATTERN 
* This pattern involves a queue of work to be done
* and multiple comcurrent workers pulling items off
* the queue
*/

package main

import (
	"fmt"
)

func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	go worker(jobs, results)

	// go worker(jobs, results) // Add more workers to take advantage of multiple core processors for faster execution
	

	for i := 0; i < 100; i++ {
		jobs <- i
	}

	close(jobs)

	for j := 0; j < 100; j++ {
		fmt.Println(<-results)
	}
}

func worker(jobs <-chan int, results chan<- int) {
	for n := range jobs {
		results <- fib(n)
	}
}

func fib(n int) int {
	if n <= 1 {
		return n
	}

	return fib(n - 1) + fib(n - 2)
}