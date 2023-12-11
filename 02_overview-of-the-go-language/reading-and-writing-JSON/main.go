package main

import (
	"encoding/json"
	"log"
)

type Person struct {
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	HairColor string `json:"hair_color"`
	HasDog bool `json:"has_dog"`
}

func main()  {
	log.Println("Hello, JSON!")

	myJSON := `
	[
		{
			"first_name": "Clark",
			"last_name": "Kent",
			"hair_color": "black",
			"has_dog": true
		},
		{
			"first_name": "Bruce",
			"last_name": "Wayne",
			"hair_color": "black",
			"has_dog": false	
		}
	]`

	var unmarshalled []Person

	err := json.Unmarshal([]byte(myJSON), &unmarshalled)
	if err != nil {
		log.Println("Error unmarshalling JSON", err)
	}

	log.Println("unmarshalled:", unmarshalled)

	//write JSON from a struct
	var mySlice []Person

	var m1 Person
	m1.FirstName = "Wally"
	m1.LastName = "West"
	m1.HairColor = "Red"
	m1.HasDog = false

	mySlice = append(mySlice, m1)

	var m2 Person
	m2.FirstName = "Deborah"
	m2.LastName = "Collins"
	m2.HairColor = "Brown"
	m2.HasDog = true

	mySlice = append(mySlice, m2)

	newJSON, err := json.MarshalIndent(mySlice, "", "     ")
	if err != nil {
		log.Println("Error marshalling slice:", err)
	} else {
		log.Println("Marshalling success.")
	}
	log.Println(string(newJSON))
}