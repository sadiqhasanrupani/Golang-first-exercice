package main

import "fmt"

func main() {
	var username string = "Sadiqhasan Rupani"

	fmt.Printf("\nMy name is: %s \nAnd The variable '%s' type is '%T'. ", username, username, username)

	var isLoggedIn bool = true
	fmt.Printf("\nVariable of this type '%t' is '%T'.", isLoggedIn, isLoggedIn)
}
