package main

import "log"

func main() {
	myVar := "bird"
	switch myVar {
	case "cat":
		log.Println("cat is set to cat")
	case "dog":
		log.Println("cat is set to dog")
	case "fish":
		log.Println("cat is set to fish")
	default:
		log.Println("cat is something else")
	}

	// myNum := 100
	// isTrue := false

	// if myNum > 99 && !isTrue {
	// 	log.Println("myNum is greater than 99 and iTrue is set to true")
	// } else if myNum < 100 && isTrue {
	// 	log.Println("1")
	// } else if myNum == 101 || isTrue {
	// 	log.Println("2")
	// } else if myNum > 1000 && !isTrue {
	// 	log.Println("3")
	// }

	// cat := "cat"

	// if cat == "cat" {
	// 	log.Println("Cat is cat")
	// } else {
	// 	log.Println("Cat is not cat")
	// }

	// isTrue := false

	// if isTrue {
	// 	log.Println("isTrue is", isTrue)
	// } else {
	// 	log.Println("isTrue is", isTrue)
	// }

}
