package main

import (
	"fmt"

	ageop "example.com/with-pointer/ageOp"
)

func main() {
	age := 32

	ageop.SubstractAge(&age, 10)
	ageAddition := ageop.AddAge(&age, 20)

	fmt.Println("Age: ", age)
	fmt.Println("Age Addition: ", ageAddition)
}
