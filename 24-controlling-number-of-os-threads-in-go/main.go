package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Number of OS threads: ", runtime.GOMAXPROCS(0))
}