package deliver_test

import (
	"pizzaDelivery/internal/deliver"
	"pizzaDelivery/internal/schemes"
	"testing"
	"github.com/stretchr/testify/require"
)

func TestPizzaDelivery(t *testing.T) {
	var m = make(map[schemes.Address]uint64)
	t.Run("Correct Output", func(t *testing.T) {
		err := deliver.PizzaDelivery("^^^^^", m)
		require.NoError(t, err)
	})

	t.Run("Correct Output", func(t *testing.T) {
		err := deliver.PizzaDelivery("^>V<^>v<", m)
		require.NoError(t, err)
	})

	t.Run("Correct Output", func(t *testing.T) {
		err := deliver.PizzaDelivery("^", m)
		require.NoError(t, err)
	})

	t.Run("Empty Input", func(t *testing.T) {
		err := deliver.PizzaDelivery("", m)
		require.NoError(t, err)
	})

	t.Run("Invalid Input", func(t *testing.T) {
		err := deliver.PizzaDelivery("abcdefg", m)
		require.Error(t, err)
	})

	t.Run("Invalid Input", func(t *testing.T) {
		err := deliver.PizzaDelivery("^<>b", m)
		require.Error(t, err)
	})
}

func TestDeliveryRouter(t *testing.T) {
	t.Run("Correct Output", func(t *testing.T) {
		var expectedOutput uint64 = 6
		res, err := deliver.DeliveryRouter(1, "^>^<^v>v")
		require.NoError(t, err)
		require.Equal(t, expectedOutput, res)
	})

	t.Run("Correct Output", func(t *testing.T) {
		var expectedOutput uint64 = 10
		res, err := deliver.DeliveryRouter(2, "^v<>^v<>^")
		require.NoError(t, err)
		require.Equal(t, expectedOutput, res)
	})

	t.Run("Correct Output", func(t *testing.T) {
		var expectedOutput uint64 = 7
		res, err := deliver.DeliveryRouter(3, "^^>Vv><V>")
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
		var expectedOutput uint64 = 1
		res, err := deliver.DeliveryRouter(4, "abcd")
		require.Error(t, err)
		require.Equal(t, expectedOutput, res)
	})

	t.Run("Invalid Input", func(t *testing.T) {
		var expectedOutput uint64 = 2
		res, err := deliver.DeliveryRouter(4, "^>v<abcd")
		require.Error(t, err)
		require.Equal(t, expectedOutput, res)
	})
	t.Run("Invalid Input", func(t *testing.T) {
		var expectedOutput uint64 = 5
		res, err := deliver.DeliveryRouter(3, "^<b>cdvef<gh")
		require.Error(t, err)
		require.Equal(t, expectedOutput, res)
	})
}