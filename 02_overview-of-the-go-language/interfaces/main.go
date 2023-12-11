package main

import "log"

type Animal interface {
	//there's must be a list of functions that any variable
	//that satisfies the animal type must have
	Says() string
	NumberOfLegs() int
}

type Dog struct {
	Name string
	Breed string
}

type Gorilla struct {
	Name string
	Colour string
	NumberOfTeeth int 
}

func main() {
	log.Println("Hello, Interfaces!")

	dog := Dog {
		Name: "Samson",
		Breed: "German Shephered",
	}

	PrintInfo(&dog)

	gorilla := Gorilla {
		Name: "Jock",
		Colour: "grey",
		NumberOfTeeth: 38,
	}

	PrintInfo(&gorilla)
}

func PrintInfo(a Animal) {
	log.Println("This animal says", a.Says(), "and has", a.NumberOfLegs(), "legs.")
}

func (d *Dog) Says() string {
	return "Woof"
}

func (d *Dog) NumberOfLegs() int {
	return 4
}

func (g *Gorilla) Says() string {
	return "Uh uh uh"
}

func (g *Gorilla) NumberOfLegs() int {
	return 2
}