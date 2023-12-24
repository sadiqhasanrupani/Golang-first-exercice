package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	revenue, err := getUserInput("Revenue: ")

	if err != nil {
		panic(err)
	}

	expenses, err := getUserInput("Expenses: ")

	if err != nil {
		panic(err)
	}

	taxRate, err := getUserInput("Tax Rate: ")

	if err != nil {
		panic(err)
	}

	ebt, profit, ratio := calculateFinancial(revenue, expenses, taxRate)
	storeResults(ebt, profit, ratio)

	fmt.Printf("%.1f\n", ebt)
	fmt.Printf("%.1f\n", profit)
	fmt.Printf("%.3f\n", ratio)
}

func storeResults(ebt, profit, ratio float64) {
	result := fmt.Sprintf("EBT: %.2f\nProfit: %.2f\nRatio: %.2f", ebt, profit, ratio)

	os.WriteFile("results.txt", []byte(result), 0644)
}

func calculateFinancial(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit
	return ebt, profit, ratio
}

func getUserInput(infoText string) (float64, error) {

	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)

	if userInput <= 0 {
		return 0, errors.New("Value must be a positive number.")
	}

	return userInput, nil
}
