package main

import (
	"fmt"
	"os"
	"log"
	"pizzaDelivery/internal/deliver"
)

func main() {
	var input string
	var deliverers int
	fmt.Println("Welcome to Thomas's Pizza Delivery Company!")
	fmt.Println("Please enter the number of deliverers")
	fmt.Scan(&deliverers)

	if len(os.Args) == 2 {
		val, err := os.ReadFile(fmt.Sprintf("routes/%s", os.Args[1]))
		if err != nil {
			log.Fatal(err)
		}
		input = string(val)
	} else {
		fmt.Println("Please enter the desired route")
		fmt.Scan(&input)
	}
	
	housesVisited, err := deliver.DeliveryRouter(deliverers, input)
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("Houses Visited = %d\n", housesVisited)
}