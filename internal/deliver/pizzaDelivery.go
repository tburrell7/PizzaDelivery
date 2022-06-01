package deliver

import (
	"errors"
	"fmt"
	"pizzaDelivery/internal/schemes"
)

//Types of errors
var OutOfBoundsError = errors.New("out of bounds error")
var InvalidInputError = errors.New("invalid input error")

//Returns number of houses visited by all deliverers and any errors
func DeliveryRouter(deliverers int, input string) (uint64, error) {
	var res uint64
	var err error
	m := make(map[schemes.Address]uint64)
	for i := 0; i < deliverers; i += 1 {
		var s string = ""
		for j, c := range input {
			if j % deliverers == i {
				s = fmt.Sprintf("%s%s", s, string(c))
			}
		}
		err = PizzaDelivery(s, m)
		if err != nil {
			res = uint64(len(m))
			return res, err
		}
	}
	res = uint64(len(m))
	return res, err
}

//Returns any errors from a deliverer
func PizzaDelivery(input string, m map[schemes.Address]uint64) error {
	//Instantiate variables
	var err error
	var x uint64 = uint64(0x8000000000000000)
	var y uint64 = uint64(0x8000000000000000)
	m[schemes.Address{X: x, Y: y}] += 1

	//Iterate through input
	for i, c := range input {
		//Validate each character
		//If out of bounds, return OutOfBoundsError
		switch string(c) {
		case ">":
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
			//if y == uint64(math.Exp2(64))-1 {
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
		m[schemes.Address{X: x, Y: y}] += 1
	}
	return err
}