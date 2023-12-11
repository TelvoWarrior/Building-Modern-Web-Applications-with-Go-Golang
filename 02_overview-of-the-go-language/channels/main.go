package main

import (
	"log"
	"example.com/helpers"
)

const numPool = 10

func CalculateValue(intChan chan int) {
	randomNumber := helpers.RandomNumber(numPool)
	intChan <- randomNumber
}

func main() {
	log.Println("Hello, Channels!")

	intChan := make(chan int)
	defer close(intChan)

	go CalculateValue(intChan)

	num := <- intChan
	log.Println(num)
}