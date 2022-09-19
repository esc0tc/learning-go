package main

import (
	"log"
	"time"
)

type User struct {
	FirstName   string
	LastName    string
	PhoneNumber string
	Age         int
	BirthDate   time.Time
}

func main() {
	user := User{
		FirstName:   "Eric",
		LastName:    "Chandler",
		PhoneNumber: "5555555555",
		Age:         100,
		BirthDate:   time.Date(1988, time.October, 12, 0, 0, 0, 0, time.UTC),
	}

	log.Println(user.FirstName, user.LastName, user.BirthDate)
}
