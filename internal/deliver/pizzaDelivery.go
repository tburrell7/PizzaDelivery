package deliver

import (
	"errors"
	"fmt"
	"math"
	"pizzaDelivery/internal/schemes"
)

//Types of errors
var OutOfBoundsError = errors.New("out of bounds error")
var InvalidInputError = errors.New("invalid input error")

//Returns number of houses visited and any errors
func PizzaDelivery(input string) (uint64, error) {
	//Instantiate variables
	var err error
	var res uint64
	var x uint64 = uint64(math.Exp2(63))
	var y uint64 = uint64(math.Exp2(63))
	m := make(map[schemes.Address]uint64)
	m[schemes.Address{X: x, Y: y}] += 1

	//Iterate through input
	for i, c := range input {
		//Validate each character
		switch string(c) {
		case "^":
			y += 1
		case "<":
			x -= 1
		case ">":
			x += 1
		case "v", "V":
			y -= 1
		//If invalid input, return InvalidInputError
		default:
			res = uint64(len(m))
			err = fmt.Errorf("%w at index %d", InvalidInputError, i)
			return res, err
		}
		//If out of bounds, return OutOfBoundsError
		if x == 0 || y == 0 || x == uint64(math.Exp2(64))-1 || y == uint64(math.Exp2(64))-1 {
			err = fmt.Errorf("%w at index %d", OutOfBoundsError, i)
			return res, err
		}
		//Increment number of deliveries to an address by 1 and store in the map
		m[schemes.Address{X: x, Y: y}] += 1
	}
	//Only visited houses will exist in the map
	res = uint64(len(m))
	return res, err
}