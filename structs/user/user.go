package user

import "time"

type User struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}
