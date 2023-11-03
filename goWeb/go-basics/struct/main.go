package main

import (
	"fmt"
	"time"
)

type User struct {
	FirstName   string
	LastName    string
	PhoneNumber string
	Address     string
	BirthDate   time.Time
	Age         int
}

func main() {
	user := User{
		FirstName: "Sadiqhasan",
		LastName:  "Rupani",
		Age:       20,
	}

	user.Address = "Sumangal vastu near khoja colony, sangli."

	fmt.Println("User name is: " + user.FirstName + " " + user.LastName)
	fmt.Println("User address is: " + user.Address)
	fmt.Println(user.FirstName + " birth date is: " + user.BirthDate.Local().Format("DD MM YYY"))
	fmt.Printf("%s %s age is \"%d\"", user.FirstName, user.LastName, user.Age)
}
