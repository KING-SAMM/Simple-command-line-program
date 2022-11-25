package main

import "fmt"

func main() {
	myBill := newBill("KC Samm's Bill")

	myBill.addItem("nodejs", 200000)
	myBill.addItem("typescript", 250000)
	myBill.addItem("reactjs", 150000)

	myBill.updateTip(10)

	fmt.Println(myBill.format())
}