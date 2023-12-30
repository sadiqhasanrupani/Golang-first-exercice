package main

import (
	"fmt"
	"time"

	"example.com/structs/utils"
)

type user struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

func (u user) outputDetails() {
	fmt.Println("\nYou First Name: ", u.firstName)
	fmt.Println("You Last Name: ", u.lastName)
	fmt.Println("You Birth date: ", u.birthDate)
}

func (u *user) clearUserName() {
	u.firstName = ""
	u.lastName = ""
}

// ^ Creating constructor
func NewUser(firstName, lastName, birthDate string) *user {
	return &user{
		firstName: firstName,
		lastName:  lastName,
		birthDate: birthDate,
	}
}

func main() {
	fName := utils.GetUserData("Please enter your first name: ")
	lName := utils.GetUserData("Please enter your last name: ")
	brithdate := utils.GetUserData("Please enter your birthdate (MM/DD/YYYY): ")

	userName := NewUser(fName, lName, brithdate)

	userName.outputDetails()
	userName.clearUserName()
	userName.outputDetails()
}
