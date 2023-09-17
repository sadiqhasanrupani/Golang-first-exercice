package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func addition(a string, b string) int {
	//^ converting a string into integer value.
	intA, aError := strconv.Atoi(a)
	intB, bError := strconv.Atoi(b)

	if aError != nil {
		fmt.Println(aError)
		return 0
	}

	if bError != nil {
		fmt.Println(aError)
		return 0
	}

	return intA + intB
}

func sub(a string, b string) int {
	//^ converting a string into integer value.
	intA, aError := strconv.Atoi(a)
	intB, bError := strconv.Atoi(b)

	if aError != nil {
		fmt.Println(aError)
		return 0
	}

	if bError != nil {
		fmt.Println(aError)
		return 0
	}

	return intA - intB
}

func multiply(a string, b string) int {
	//^ converting a string into integer value.
	intA, aError := strconv.Atoi(a)
	intB, bError := strconv.Atoi(b)

	if aError != nil {
		fmt.Println(aError)
		return 0
	}

	if bError != nil {
		fmt.Println(aError)
		return 0
	}

	return intA * intB
}

func division(a string, b string) float64 {
	//^ converting a string into integer value.
	intA, aError := strconv.ParseFloat(a, 64)
	intB, bError := strconv.ParseFloat(b, 64)

	if aError != nil {
		fmt.Println(aError)
		return 0
	}

	if bError != nil {
		fmt.Println(aError)
		return 0
	}

	return intA / intB
}

func main() {
	//^ arithmetic operation

	//^ getting value from the users
	valueA := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the value of a: ")

	inputA, _ := valueA.ReadString('\n')
	inputA = strings.TrimSpace(inputA)

	valueB := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the value of b: ")

	inputB, _ := valueB.ReadString('\n')
	inputB = strings.TrimSpace(inputB)

	sumResult := addition(inputA, inputB)
	subResult := sub(inputA, inputB)
	multiplyResult := multiply(inputA, inputB)
	divisionResult := division(inputA, inputB)

	fmt.Printf("\nValue of %s + %s is %d \n", inputA, inputB, sumResult)
	fmt.Printf("\nValue of %s - %s is %d \n", inputA, inputB, subResult)
	fmt.Printf("\nValue of %s * %s is %d \n", inputA, inputB, multiplyResult)
	fmt.Printf("\nValue of %s / %s is %f \n", inputA, inputB, divisionResult)

}
