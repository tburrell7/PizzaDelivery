package deliver_test

import (
	"pizzaDelivery/internal/deliver"
	"pizzaDelivery/internal/schemes"
	"testing"

	"github.com/stretchr/testify/require"
)
func clearMap(m map[schemes.Address]uint64) {
	for a := range m {
		delete(m, a)
	}
}
func TestPizzaDelivery(t *testing.T) {
	var m = make(map[schemes.Address]uint64)
	t.Run("Correct Output", func(t *testing.T) {
		err := deliver.PizzaDelivery("^^^^^", m)
		require.NoError(t, err)
	})

	clearMap(m)
	t.Run("Correct Output", func(t *testing.T) {
		err := deliver.PizzaDelivery("^>V<", m)
		require.NoError(t, err)
	})

	clearMap(m)
	t.Run("Correct Output", func(t *testing.T) {
		err := deliver.PizzaDelivery("^", m)
		require.NoError(t, err)
	})

	clearMap(m)
	t.Run("Empty Input", func(t *testing.T) {
		err := deliver.PizzaDelivery("", m)
		require.NoError(t, err)
	})

	clearMap(m)
	t.Run("Invalid Input", func(t *testing.T) {
		err := deliver.PizzaDelivery("abcdefg", m)
		require.Error(t, err)
	})

	clearMap(m)
	t.Run("Invalid Input", func(t *testing.T) {
		err := deliver.PizzaDelivery("^<>b", m)
		require.Error(t, err)
	})
}

func TestDeliveryRouter(t *testing.T) {
	t.Run("Correct Output", func(t *testing.T) {
		var expectedOutput uint64 = 102
		res, err := deliver.DeliveryRouter(1, "^^<<v<<v><v^^<><>^^<v<v^>>^^^><^>v^>v><><><<vv^^<^>^^<v^>v>v^v>>>^<>v<^<v^><^>>>>><<v>>^>>^>v^>><<^>v>v<>^v^v^vvv><>^^>v><v<><>^><^^<vv^v<v>^v>>^v^>v><>v^<vv>^><<v^>vv^<<>v>>><<<>>^<vv<^<>^^vv>>>^><<<<vv^v^>>><><^>v<>^>v<v^v<^vv><^v^><<<<>^<>v>^v>v<v<v<<>")
		require.NoError(t, err)
		require.Equal(t, expectedOutput, res)
	})
	t.Run("Correct Output", func(t *testing.T) {
		var expectedOutput uint64 = 4
		res, err := deliver.DeliveryRouter(2, "^^^^^^")
		require.NoError(t, err)
		require.Equal(t, expectedOutput, res)
	})

	t.Run("Correct Output", func(t *testing.T) {
		var expectedOutput uint64 = 3
		res, err := deliver.DeliveryRouter(3, "^^^^^^")
		require.NoError(t, err)
		require.Equal(t, expectedOutput, res)
	})

	t.Run("Empty Input", func(t *testing.T) {
		var expectedOutput uint64 = 1
		res, err := deliver.DeliveryRouter(5, "")
		require.NoError(t, err)
		require.Equal(t, expectedOutput, res)
	})

	t.Run("Invalid Input", func(t *testing.T) {
		_, err := deliver.DeliveryRouter(4, "abcd")
		require.Error(t, err)
	})
}