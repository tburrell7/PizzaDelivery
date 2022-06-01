package main

import (
	"fmt"
	"log"
	"pizzaDelivery/internal/deliver"
)

func main() {
	var input string
	var deliverers int
	fmt.Println("Welcome to Thomas's Pizza Delivery Company!")
	fmt.Println("Please enter the number of deliverers")
	fmt.Scan(&deliverers)
	fmt.Println("Please enter the desired route")
	fmt.Scan(&input)
	//Scanner
	housesVisited, err := deliver.DeliveryRouter(deliverers, input)
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("Houses Visited = %d\n", housesVisited)
}