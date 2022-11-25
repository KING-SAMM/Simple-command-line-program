package main

import (
	"bufio"
	"os"
	"fmt"
	"strings"
)

// Receive user input
func getInput(prompt string, r *bufio.Reader) (string, error) {
	// prompt user
	fmt.Print(prompt)

	// Receive user input
	input, err := r.ReadString('\n')
	// Trim spaces
	return strings.TrimSpace(input), err
}

func createBill() bill {
	reader := bufio.NewReader(os.Stdin)

	name, _ := getInput("Create new bill, enter bill name: ", reader)

	// Create bill with name
	b := newBill(name)

	fmt.Println("Created bill - ", b.name)

	return b
}

func billOptions(b bill) {
	reader := bufio.NewReader(os.Stdin)

	option, _ := getInput("Select an option (a - add Item, s - save bill, t - add tip): ", reader)

	switch option {
	case "a":
		name, _ := getInput("Item name: ", reader)
		price, _ := getInput("Item price: ", reader)

		p, err := strconnv.ParseFloat(price, 64)
		if err != nil {
			fmt.Println("The price must be a number")
			billOptions(b)
		}
		b.addItem(name, p)

		fmt.Println("Item added - " name, price)
	case "t":
		tip, _ := getInput("Enter tip amount (N): ", reader)
		fmt.Println(tip)
		
		t, err := strconnv.ParseFloat(tip, 64)
		if err != nil {
			fmt.Println("The tip must be a number")
			billOptions(b)
		}
		b.updateTip(t)

		fmt.Println("Tip added - " tip)

		billOptions(b)

	case "s":
		fmt.Println("You chose s")
	default:
		fmt.Println("That was not a valid option")
		
		billOptions(b)
	}
}

func main() {
	mybill := createBill()
	billOptions(mybill)
}