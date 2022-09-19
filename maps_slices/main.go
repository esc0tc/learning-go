package main

import (
	"log"
)

type User struct {
	FirstName string
	LastName  string
}

func main() {

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	log.Println(numbers[0:2])

	names := []string{"one", "seven", "fish", "cat"}

	log.Println(names[0:2])
	// var mySlice []int // This is now a slice of strings

	// mySlice = append(mySlice, 2)
	// mySlice = append(mySlice, 1)
	// mySlice = append(mySlice, 3)

	// sort.Ints(mySlice)

	// log.Println(mySlice)

	// myMap := make(map[string]User)

	// me := User{
	// 	FirstName: "Eric",
	// 	LastName:  "Chandler",
	// }

	// myMap["me"] = me

	// log.Println(myMap["me"].FirstName)

	// myMap["First"] = 1
	// myMap["Second"] = 2

	// myMap := make(map[string]string)

	// myMap["dog"] = "Remi"
	// myMap["otherDog"] = "Cassy"
	// myMap["dog"] = "Fido"

	// log.Println(myMap["First"])
	// log.Println(myMap["Second"])
}
