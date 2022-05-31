package deliver

import (
	"errors"
	"fmt"
	"pizzaDelivery/internal/schemes"
)

var outOfBoundsError = errors.New("out of bounds error")
var invalidInputError = errors.New("invalid input error")

func PizzaDelivery(input string) (uint64, error) {
	//iterate through input, validating along the way. Store position and times delivered in hashmap
	//store result in hash map
	var err error
	var res uint64
	var x uint64 = 2^63
	var y uint64 = 2^63
	m := make(map[schemes.Address]uint64)
	m[schemes.Address{X: x, Y: y}] += 1

	for i, c := range input {
		switch string(c) {
		case "^":
			y += 1
		case "<":
			x -= 1
		case ">":
			x += 1
		case "v", "V":
			y -= 1
		default:
			res = uint64(len(m))
			err = fmt.Errorf("%w at index %d", invalidInputError, i)
			return res, err
		}
		m[schemes.Address{X: x, Y: y}] += 1
	}
	res = uint64(len(m))
	return res, err
}