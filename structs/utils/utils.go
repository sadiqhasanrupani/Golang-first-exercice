package utils

import (
	"fmt"

	"example.com/structs/user"
)

func GetUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}

func OutputUserName(u user.User) {
	fmt.Println(u)
}
