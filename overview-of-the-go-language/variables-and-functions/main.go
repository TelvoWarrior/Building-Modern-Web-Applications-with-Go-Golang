package main

import "fmt"

func main() {
	fmt.Println("Hello, world.")

	var whatToSay string
	var i int

	whatToSay = "Goodbye, cruel world."
	fmt.Println(whatToSay)

	i = 7
	fmt.Println("i is set to:", i)

	whatWasSaid, whatWasElseSaid := saySomething()
	fmt.Println(whatWasSaid, whatWasElseSaid)
}

func saySomething() (string, string) {
	return "something", "else"
}