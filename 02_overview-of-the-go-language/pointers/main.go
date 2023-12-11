package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("Hello, Pointers!")

	var myString string
	myString = "Green"

	//I notice the difference between fmt and log.
	//fmt just prints a line
	//log prints data, time and a line
	log.Println("myString is set to", myString)

	//we pass the link, reference to the memory location where the value lives
	//we said: we want to pass not a particular value, but a reference to the value
	changeUsingPointers(&myString)

	log.Println("myString now is", myString)
}

func changeUsingPointers(s *string) {
	//we recieve the reference where the value lives
	//we print that reference
	log.Println("s is set to:", s)
	log.Println("let's try to dereference the pointer:", *s)
	newValue := "Red"
	//we change the value at the memory location
	*s = newValue
	//function return nothing
	//it changes the value in the memory location
}