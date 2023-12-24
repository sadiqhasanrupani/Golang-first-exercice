package main

import (
	"fmt"
	"math"
)

const inflationRate = 6.5

func main() {
	investmentAmount := getInput("Enter your investment amount: ")
	expectedReturnRate := getInput("\nEnter the expected return date:")
	years := getInput("\nEnter your no. of years: ")

	futureValue, futureRealValue := calculateFutureValues(investmentAmount, expectedReturnRate, years)

	futureValueOutput := fmt.Sprintf("\nThe future value is: %.2f", futureValue)
	futureRealValOutput := fmt.Sprintf("\nThe Future real value is: %2.f", futureRealValue)

	fmt.Print(futureValueOutput, futureRealValOutput)
}

func outputText(text string) {
	fmt.Print(text)
}

func getInput(message string) float64 {
	var userInput float64
	outputText(message)
	fmt.Scan(&userInput)
	return userInput
}

func calculateFutureValues(investmentAmount, expectedReturnRate, years float64) (fv float64, rfv float64) {
	fv = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	rfv = fv / math.Pow(1+inflationRate/100, years)

	// return fv, rfv OR
	return
}

// func calculateFutureValues(investmentAmount, expectedReturnRate, years float64) (float64, float64) {
// 	fv := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
// 	rfv := fv / math.Pow(1+inflationRate/100, years)

// 	return fv, rfv
// }
