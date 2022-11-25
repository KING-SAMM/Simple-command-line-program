package main

import "fmt"

func main() {
	myBill := newBill("KC Samm's Bill")

	fmt.Println(myBill.format())
}