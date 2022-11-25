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
	fmt.Println(option)
}

func main() {
	mybill := createBill()
	billOptions(mybill)
}