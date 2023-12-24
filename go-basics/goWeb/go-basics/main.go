package main

import "fmt"

func main() {
	var something string = "This is something"
	var someIntValue int
	someIntValue = 90

	fmt.Println("Their is string saying \"" + something + "\"")
	fmt.Println("Their is some integer value also which is", someIntValue)

	somethingFromFunc := returnString()

	fmt.Println("This function returns: ", somethingFromFunc)
}

//^ Creating another function
func returnString() string {
	return "Something string coming from the return string function"
}
