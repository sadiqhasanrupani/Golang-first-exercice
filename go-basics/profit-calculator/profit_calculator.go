package main

import "fmt"

func main() {
	// Asking the user for revenue, expenses and tax rate
	fmt.Print("Enter revenue: ")
	var revenue float64
	fmt.Scan(&revenue)

	fmt.Print("Enter expenses: ")
	var expenses float64
	fmt.Scan(&expenses)

	fmt.Print("Enter tax rate: ")
	var taxRate float64
	fmt.Scan(&taxRate)

	earningsBeforeTax := revenue - expenses
	earningsAfterTax := earningsBeforeTax - taxRate

	ratio := earningsBeforeTax / earningsAfterTax

	// Outputting the desired ones
	fmt.Printf("\nEBT: %.2f", earningsBeforeTax)
	fmt.Printf("\nProfit: %.2f", earningsAfterTax)
	fmt.Printf("\nRatio: %.2f", ratio)
}
