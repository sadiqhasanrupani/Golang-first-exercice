package main

import "fmt"

//^ Creating a basic struct
type User struct {
	FullName string
	Age      int
	Address  string
}

func main() {
	//^ Creating a basic map like this
	basicStringMap := make(map[string]string)

	//^ Creating a map and adding a User struct in it
	userStructMap := make(map[string]User)
	user1 := User{
		FullName: "John Doe",
		Age:      30,
		Address:  "New York City",
	}

	userStructMap["user1"] = user1

	basicStringMap["name"] = "Sadiqhasan Rupani"
	basicStringMap["Age"] = "20"

	fmt.Printf("Name: %s and Age: %s", basicStringMap["name"], basicStringMap["Age"])
	fmt.Printf("\nSo the full name of user1 is: %s and his/her age is: %d\n", userStructMap["user1"].FullName, userStructMap["user1"].Age)
	fmt.Println(userStructMap)
}
