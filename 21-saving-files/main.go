package main

import (
	"bufio"
	"os"
	"fmt"
	"strings"
	"strconv"
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

		// Convert user input (string) to float
		p, err := strconv.ParseFloat(price, 64)
		if err != nil {
			fmt.Println("The price must be a number")
			billOptions(b)
		}
		b.addItem(name, p)

		fmt.Println("Item added - ", name, price)
		billOptions(b)

	case "t":
		tip, _ := getInput("Enter tip amount (N): ", reader)
		fmt.Println(tip)
		
		t, err := strconv.ParseFloat(tip, 64)
		if err != nil {
			fmt.Println("The tip must be a number")
			billOptions(b)
		}
		b.updateTip(t)

		fmt.Println("Tip added - ", tip)
		billOptions(b)

	case "s":
		b.save()

		fmt.Printf("Bill, '%v', saved successfully \n", b.name)
	default:
		fmt.Println("That was not a valid option")
		
		billOptions(b)
	}
}

func main() {
	mybill := createBill()
	billOptions(mybill)
}