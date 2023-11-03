package main

import (
	"fmt"
	"time"
)

type User struct {
	Name      string
	Age       int
	BirthDate time.Time
}

func main() {

	var user User
	user.Name = "Sadiqhasan Rupani"

	fmt.Printf("The User struct field names: %+v and values: %#v", user, user)
}
