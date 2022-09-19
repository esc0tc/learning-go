package main

import (
	"log"

	"github.com/esc0tc/go_package_1/helpers"
)

func main() {
	log.Println("hello")
	var myVar helpers.SomeType
	myVar.TypeName = "Some Name"
	log.Println(myVar.TypeName)
}
