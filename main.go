package main

import (
	"fmt"
	"log"
	"pizzaDelivery/internal/deliver"
)

func main() {
	var input string
	fmt.Println("Welcome to Thomas's Pizza Delivery Company!")
	fmt.Println("Please enter the desired route")
	fmt.Scan(&input)
	housesVisited, err := deliver.PizzaDelivery(input)
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("Houses Visited = %d\n", housesVisited)
}