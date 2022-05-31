package main

import (
	"fmt"
	"log"
	"pizzaDelivery/internal/deliver"
)

func main() {
	housesVisited, err := deliver.PizzaDelivery("^^f^")
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("Houses Visited = %d\n", housesVisited)
}