package main

import "fmt"

// Every go file must end in .go
// Every go file must start with a package declaration
// Every go program has to have a function called main

func main() {
	fmt.Println("Hello world.")

	var whatToSay string
	var i int

	whatToSay = "Goodbye, cruel world"
	fmt.Println(whatToSay)

	i = 7

	fmt.Println("i is set to", i)

	whatWasSaid, theOtherThingThatWasSaid := saySomething()

	fmt.Println(whatWasSaid, theOtherThingThatWasSaid)
}

func saySomething() (string, string) {
	return "something", "else"
}
