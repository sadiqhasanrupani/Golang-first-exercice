package main

import (
	"fmt"

	"example.com/structs/user"
	"example.com/structs/utils"
)

func main() {
	fName := utils.GetUserData("Please enter your first name: ")
	lName := utils.GetUserData("Please enter your last name: ")
	brithdate := utils.GetUserData("Please enter your birthdate (MM/DD/YYYY): ")

	userName, err := user.NewUser(fName, lName, brithdate)

	if err != nil {
		fmt.Println(err)
		return
	}

	userName.OutputDetails()
	userName.ClearUserName()
	userName.OutputDetails()
}
