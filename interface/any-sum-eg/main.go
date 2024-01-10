package main

import "fmt"

func main() {
	value := sumOfAny(2, 45)
	fmt.Println(value)
}

func sumOfAny[T string | float64 | int](a, b T) T {
	return a + b
}
