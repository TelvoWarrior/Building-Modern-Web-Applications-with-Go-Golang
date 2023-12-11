package main

import (
	"github.com/TelvoWarrior/Building-Modern-Web-Applications-with-Go-Golang/tree/main/02_overview-of-the-go-language/packages/helpers"
	"log"
)

func main() {
	log.Println("Hello, Packages!")

	var myVar helpers.SomeType
	myVar.TypeName = "This is type name"
	log.Println(myVar)

}