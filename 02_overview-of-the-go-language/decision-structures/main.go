package main

import "log"

func main() {
	log.Println("Hello, Decision Structures!")

	var isTrue bool
	isTrue = false

	if isTrue {
		log.Println("isTrue is:", isTrue)
	} else {
		log.Println("isTrue is:", isTrue)
	}

	myVar := "zergling"

	switch myVar {
	case "cat":
		log.Println("It's a cat.")
	case "dog":
		log.Println("It's a dog.")
	case "fish":
		log.Println("It's a fish.")
	default:
		log.Println("It's an unknown creature.")
	}
}