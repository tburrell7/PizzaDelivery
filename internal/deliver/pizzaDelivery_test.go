package deliver_test

import (
	"pizzaDelivery/internal/deliver"
	"testing"
	"github.com/stretchr/testify/require"
)

func TestPizzaDelivery(t *testing.T) {
	t.Run("Correct Output", func(t *testing.T) {
		var expectedOutput uint64 = 6
		res, err := deliver.PizzaDelivery("^^^^^")
		require.NoError(t, err)
		require.Equal(t, res, expectedOutput)
	})
	t.Run("Correct Output", func(t *testing.T) {
		var expectedOutput uint64 = 4
		res, err := deliver.PizzaDelivery("^>V<")
		require.NoError(t, err)
		require.Equal(t, res, expectedOutput)
	})
	t.Run("Correct Output", func(t *testing.T) {
		var expectedOutput uint64 = 2
		res, err := deliver.PizzaDelivery("^")
		require.NoError(t, err)
		require.Equal(t, res, expectedOutput)
	})
	t.Run("Empty Input", func(t *testing.T) {
		var expectedOutput uint64 = 1
		res, err := deliver.PizzaDelivery("")
		require.NoError(t, err)
		require.Equal(t, res, expectedOutput)
	})
	t.Run("Invalid Input", func(t *testing.T) {
		_, err := deliver.PizzaDelivery("abcdefg")
		require.Error(t, err)
	})
	t.Run("Invalid Input", func(t *testing.T) {
		_, err := deliver.PizzaDelivery("^<>b")
		require.Error(t, err)
	})
}