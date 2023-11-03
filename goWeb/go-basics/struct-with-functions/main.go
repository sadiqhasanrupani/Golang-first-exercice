package main

import (
	"fmt"
	"time"
)

// ^ Creating a simple struct of user
type User struct {
	FirstName string
	LastName  string
	Address   string
	Age       int
	BirthDate time.Time
}

// ^ Creating a function which will be connected the simple user struct
func (m *User) displayFullName() string {
	return m.FirstName + " " + m.LastName
}

func main() {
	//^ creating to users and give them a value of first-name and last-name
	var user1 User
	user1.FirstName = "Sadiqhasan"
	user1.LastName = "Rupani"

	user2 := User{
		FirstName: "Aliabbas",
		LastName:  "Rupani",
	}

	fmt.Printf("User no.1 full name is: \"%s\"\n", user1.displayFullName())
	fmt.Printf("User no.2 full name is: \"%s\"\n", user2.displayFullName())
}
