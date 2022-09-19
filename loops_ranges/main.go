package main

import "log"

func main() {

	type User struct {
		FirstName string
		LastName  string
		Email     string
		Age       int
	}

	var users []User
	users = append(users, User{"John", "Smith", "john@smith.com", 30})
	users = append(users, User{"Mary", "Jones", "mary@jones.com", 25})
	users = append(users, User{"Sally", "Smith", "sally@smith.com", 45})
	users = append(users, User{"Alex", "Smith", "alex@smith.com", 17})

	// var firstLine = "once upon a time"

	for _, l := range users {
		log.Println(l.FirstName, l.LastName, l.Email, l.Age)
	}

	// animals := []string{"dog", "fish", "horse", "cow", "cat"}
	// animals := make(map[string]string)
	// animals["dog"] = "fido"
	// animals["cat"] = "stripes"
	// for animalType, animal := range animals {
	// 	log.Println(animalType, animal)
	// }
}
