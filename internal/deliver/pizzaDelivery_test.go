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
		var expectedOutput uint64 = 6
		res, err := deliver.PizzaDelivery("^^^^^", m)
		require.NoError(t, err)
		require.Equal(t, res, expectedOutput)
	})

	clearMap(m)
	t.Run("Correct Output", func(t *testing.T) {
		var expectedOutput uint64 = 4
		res, err := deliver.PizzaDelivery("^>V<", m)
		require.NoError(t, err)
		require.Equal(t, res, expectedOutput)
	})

	clearMap(m)
	t.Run("Correct Output", func(t *testing.T) {
		var expectedOutput uint64 = 2
		res, err := deliver.PizzaDelivery("^", m)
		require.NoError(t, err)
		require.Equal(t, res, expectedOutput)
	})

	clearMap(m)
	t.Run("Empty Input", func(t *testing.T) {
		var expectedOutput uint64 = 1
		res, err := deliver.PizzaDelivery("", m)
		require.NoError(t, err)
		require.Equal(t, res, expectedOutput)
	})

	clearMap(m)
	t.Run("Invalid Input", func(t *testing.T) {
		_, err := deliver.PizzaDelivery("abcdefg", m)
		require.Error(t, err)
	})

	clearMap(m)
	t.Run("Invalid Input", func(t *testing.T) {
		_, err := deliver.PizzaDelivery("^<>b", m)
		require.Error(t, err)
	})
}

func TestDeliveryRouter(t *testing.T) {
	t.Run("Correct Output", func(t *testing.T) {
		var expectedOutput uint64 = 4
		res, err := deliver.DeliveryRouter(2, "^^^^^^")
		require.NoError(t, err)
		require.Equal(t, res, expectedOutput)
	})

	t.Run("Correct Output", func(t *testing.T) {
		var expectedOutput uint64 = 3
		res, err := deliver.DeliveryRouter(3, "^^^^^^")
		require.NoError(t, err)
		require.Equal(t, res, expectedOutput)
	})

	t.Run("Empty Input", func(t *testing.T) {
		var expectedOutput uint64 = 1
		res, err := deliver.DeliveryRouter(5, "")
		require.NoError(t, err)
		require.Equal(t, res, expectedOutput)
	})

	t.Run("Invalid Input", func(t *testing.T) {
		_, err := deliver.DeliveryRouter(4, "abcd")
		require.Error(t, err)
	})
}