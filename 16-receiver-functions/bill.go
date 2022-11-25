package main

import "fmt"

type bill struct {
	name string
	items map[string]float64
	tip float64
}

// Make new bills
func newBill(name string) bill {
	b := bill{
		name: name,
		items: map[string]float64{"solidity": 234000, "golang": 205000, "javascript": 120000},
		tip: 0,
	}

	return b
}

// Receiver function to format the bill
func (b bill) format() string {
	fs := "Course Breakdown: \n"
	var total float64 = 0

	// List out the items
	for k, v := range b.items {
		fs += fmt.Sprintf("%-12v ---- N%v \n", k+":", v)
		total += v
	}

	// Total
	fs += fmt.Sprintf("%-12v ---- N%0.2f \n", "total", total)

	return fs

}