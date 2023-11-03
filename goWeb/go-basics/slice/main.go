package main

import "fmt"

/*
^ To create a slice you need to define like this,
! var datatype []string
*/

func main() {
	var sliceString []string

	//^ To add a a string or any other datatype inside a slice you use append method like this
	sliceString = append(sliceString, "Sadiqhasan Rupani")

	//^ You can also create a slice using a shorthand property
	numberSlice := []int{1, 2, 3, 4, 5, 6, 7}

	fmt.Printf("The number slice are %v", numberSlice)
	fmt.Printf("\nThe string slice are %v", sliceString)

	//^ you can see range of item inside a slice using this syntax
	fmt.Printf("\nThe range between 0 and 4 of number slice is: %v", numberSlice[0:4])
	fmt.Printf("\nThe 2nd index value slice is: %v", numberSlice[2])
}
