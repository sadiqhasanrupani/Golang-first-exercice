package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	FirstName string
	LastName  string
	BirthDate string
	CreatedAt time.Time
}

func (u User) OutputDetails() {
	fmt.Println("\nYou First Name: ", u.FirstName)
	fmt.Println("You Last Name: ", u.LastName)
	fmt.Println("You Birth date: ", u.BirthDate)
}

func (u *User) ClearUserName() {
	u.FirstName = ""
	u.LastName = ""
}

// ^ Creating constructor
func NewUser(firstName, lastName, birthDate string) (*User, error) {

	if firstName == "" || lastName == "" || birthDate == "" {
		return nil, errors.New("firstName, lastName and birthDate are required")
	}

	return &User{
		FirstName: firstName,
		LastName:  lastName,
		BirthDate: birthDate,
		CreatedAt: time.Now(),
	}, nil
}
