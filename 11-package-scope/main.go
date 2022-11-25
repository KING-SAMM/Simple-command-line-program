package main

import "fmt"

/* 
mensuration.go variables and functions are automatically 
available here without being explicitly 
imported, so long as they belong to the same "main" 
package, and both files are compiled and run
*/

func main() {
	c := circumference(3.2)
	v := vol_of_cuboid(2, 4, 6)

	fmt.Printf("Circumference = %0.1f \nVolume of cuboid = %0.1f", c, v)
}