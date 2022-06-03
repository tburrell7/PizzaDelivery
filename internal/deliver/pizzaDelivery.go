package deliver

import (
	"errors"
	"fmt"
)
//public for testing purposes
type Address struct {
	x uint64
	y uint64
}

//ORIGIN = 2^63
const ORIGIN = uint64(0x8000000000000000)

//Types of errors
var OutOfBoundsError = errors.New("out of bounds error")
var InvalidInputError = errors.New("invalid input error")

//Returns number of houses visited by all deliverers and any errors
func DeliveryRouter(deliverers int, input string) (uint64, error) {
	var res uint64
	var err error
	//All deliverers are in the same neighborhood so they access the same map
	neighborhood := make(map[Address]uint64)
	//Parse input based on number of deliverers
	for deliverer := 0; deliverer < deliverers; deliverer += 1 {
		var delRoute string = ""
		for index, c := range input {
			direction := string(c)
			if index % deliverers == deliverer {
				delRoute = fmt.Sprintf("%s%s", delRoute, direction)
			}
		}
		err = PizzaDelivery(delRoute, neighborhood)
		if err != nil {
			res = uint64(len(neighborhood))
			return res, err
		}
	}
	//Only visited addresses are added to m
	res = uint64(len(neighborhood))
	return res, err
}

//Returns any errors from a deliverer
func PizzaDelivery(input string, neighborhood map[Address]uint64) error {
	//Instantiate variables
	var err error
	var x uint64 = ORIGIN
	var y uint64 = ORIGIN
	neighborhood[Address{x, y}] += 1

	//Iterate through input
	for i, c := range input {
		//Validate each character
		//If out of bounds, return OutOfBoundsError
		switch string(c) {
		case ">":
			//if x == 2^64-1
			if x == uint64(0xffffffffffffffff) {
				err = fmt.Errorf("%w at index %d", OutOfBoundsError, i)
				return err
			}
			x += 1
		case "<":
			if x == 0 {
				err = fmt.Errorf("%w at index %d", OutOfBoundsError, i)
				return err
			}
			x -= 1
		case "^":
			//if y == 2^64-1
			if y == uint64(0xffffffffffffffff) {
				err = fmt.Errorf("%w at index %d", OutOfBoundsError, i)
				return err
			}
			y += 1
		case "v", "V":
			if y == 0 {
				err = fmt.Errorf("%w at index %d", OutOfBoundsError, i)
				return err
			}
			y -= 1
		//If invalid input, return InvalidInputError
		default:
			err = fmt.Errorf("%w at index %d", InvalidInputError, i)
			return err
		}
		//Increment number of deliveries to an address by 1 and store in the map
		neighborhood[Address{x, y}] += 1
	}
	return err
}