package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings" // Import the "strings" package for string manipulation
)

func main() {
	greetingMsg := "Welcome to the Golang journey, "

	userName := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter your name: ")

	name, _ := userName.ReadString('\n')

	fmt.Printf("\n%s%s", greetingMsg, name)

	resultOfSum := sum()

	fmt.Printf("The Value of the sum is: %d", resultOfSum)
}

func sum() int64 {
	inputA := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the value of a: ")
	aInput, _ := inputA.ReadString('\n')
	aInput = strings.TrimSpace(aInput) // Remove leading/trailing whitespaces

	inputB := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the value of b: ")
	bInput, _ := inputB.ReadString('\n')
	bInput = strings.TrimSpace(bInput) // Remove leading/trailing whitespaces

	a, valueAErr := strconv.ParseInt(aInput, 10, 64)
	b, valueBErr := strconv.ParseInt(bInput, 10, 64)

	if valueAErr != nil {
		fmt.Println("\n", valueAErr)
		return 0
	}

	if valueBErr != nil {
		fmt.Println("\n", valueBErr)
		return 0
	}

	return a + b
}
