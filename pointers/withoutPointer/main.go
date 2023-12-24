package main

import (
	"fmt"

	"example.com/without-pointer/withoutPointer/age"
)

func main() {
	//^ creating a age variable of 32
	var ageValue int = 32

	minusAge := age.DecreaseAge(ageValue, 18)

	fmt.Println("Age: ", minusAge)
}
