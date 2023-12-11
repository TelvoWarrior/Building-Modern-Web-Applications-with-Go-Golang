package main

import (
	"log"
	"time"
)

//capitalising means that it can be used in other packages
type User struct {
	FirstName string
	LastName string
	PhoneNumber string
	Age int
	BirthDate time.Time
}

func main()  {
	user := User {
		FirstName: "Trevor",
		LastName: "Sawler",
	}
	
	log.Println("user's first name:", user.FirstName)
	log.Println("user's last name:", user.LastName)
	log.Println("user's birthdate:", user.BirthDate)
}

//variable s declared outside a function so it's available to be used inside any function
// var s = "seven"

// func main() {
// 	log.Println("Hello, Types and Structs!")

// 	//variable s2 declared in function main so it's available to be used only inside function main
// 	var s2 = "six"

// 	s := "eight"

// 	log.Println("s is:", s)
// 	log.Println("s2 is:", s2)

// 	saySomething("xxx")
// }

// func saySomething(s3 string) (string, string) {
// 	log.Println("s from the saySomething function is:", s)
// 	return s3, "World"
// }