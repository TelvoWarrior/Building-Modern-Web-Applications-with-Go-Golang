package main

import "log"

func main() {
	log.Println("Hello, Loops and Ranging!")

	//in Go there's a one kind of loop - for
	for i := 0; i <= 10; i++ {
		log.Println("Current value of i is:", i)
	}

	animals := []string { "dog", "fish", "cow", "horse", "cat" }

	for i, v := range animals {
		log.Println("Index is:", i, "Value is:", v)
	}

	//you can use blank identifier if you don't care about index
	for _, v := range animals {
		log.Println("Value is:", v)
	}

	//if you need just an index then you use one variable without blank identifier
	for i := range animals {
		log.Println("Index is:", i)
	}

	//ranging over a map
	animalsMap := make(map[string]string)
	animalsMap["dog"] = "Fido"
	animalsMap["cat"] = "Fluffy"

	for k, v := range animalsMap {
		log.Println("map key is:", k, "map value is:", v)
	}

	//ranging over a string
	var firstLine = "Once upon a midnight dreary"

	for i, l := range firstLine {
		log.Println("index:", i, "letter:", l)
		//in Go string is actually a slice of bytes
	}

	type User struct {
		FirstName string
		LastName string
		Email string
		Age int
	}

	var users []User
	users = append(users, User{ "John", "Smith", "john@smith.com", 30 })
	users = append(users, User{ "Mary", "Jones", "mary@jones.com", 25 })
	users = append(users, User{ "Deborah", "Giggs", "deborah@giggs.com", 33 })
	users = append(users, User{ "Frank", "Pete", "frank@pete.com", 27 })

	for _, u := range users {
		log.Println(u.FirstName, u.LastName, u.Email, u.Age)
	}
}