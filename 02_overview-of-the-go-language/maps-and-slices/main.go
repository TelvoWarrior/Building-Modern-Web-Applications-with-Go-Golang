package main

import (
	"log"
	"sort"
)

type User struct {
	FirstName string
	LastName string
}

func main() {

	//don't create maps that way; why? you should find it out
	// var myOtherMap map[string]string

	//create maps this way
	myMap := make(map[string]string)
	myMap["dog"] = "Samson"
	myMap["otherDog"] = "Cassie"
	myMap["dog"] = "Fido"

	log.Println(myMap["dog"])
	log.Println(myMap["otherDog"])

	mySecondMap := make(map[string]int)
	mySecondMap["First"] = 1
	mySecondMap["Second"] = 2
	log.Println(mySecondMap["First"], mySecondMap["Second"])
	
	//maps can store struct types
	//[key type]value type; key is like an index; 
	//maps are unsorted, you always must look up by key and don't rely on the order you put key-values in a map;
	myUserMap := make(map[string]User)

	me := User {
		FirstName: "Trevor",
		LastName: "Sawler",
	}

	myUserMap["me"] = me

	log.Println(myUserMap["me"].FirstName)

	//Slices section
	var mySlice []int

	//that's how you add new element to a slice
	mySlice = append(mySlice, 2)
	mySlice = append(mySlice, 1)
	mySlice = append(mySlice, 3)
	
	sort.Ints(mySlice)

	log.Println(mySlice)

	numbers := []int{ 1, 2, 3, 4, 5, 6, 7, 8, 9, 10 }
	log.Println(numbers)
	log.Println(numbers[0:2])

	words := []string{ "fish", "cat", "one", "seven", "dog" }
	log.Println(words)
}