package main

func main() {
	sumOfAny(2, 45)
}

func sumOfAny[T string | float64 | int](a, b T) T {
	return a + b
}
