package hpi_difference

import (
	"github.com/web3foru/computer-science-group-golang/internal/application/src/algorithms/simple_linked_list_applications/high_precision_integers"
	"github.com/web3foru/computer-science-group-golang/internal/application/src/algorithms/simple_linked_list_applications/high_precision_integers/difference_hpi"
	"testing"
)

func TestDifferenceHighPrecisionInteger(t *testing.T) {
	t.Run("subtract a HPI from another, both with only one node each", func(t *testing.T) {
		n1 := high_precision_integers.SetUpHighPrecisionInteger(999)
		n2 := high_precision_integers.SetUpHighPrecisionInteger(700)
		result := difference_hpi.Subtract(&n1, &n2)

		if result.NumberRepresentation.GetFirstNode().GetData().(int) != 299 {
			t.Error("Subtraction has not well performed")
		}
	})

	t.Run("subtract a HPI from another, both with only one node each, subtrahend is greater", func(t *testing.T) {
		n1 := high_precision_integers.SetUpHighPrecisionInteger(999)
		n2 := high_precision_integers.SetUpHighPrecisionInteger(1000)
		result := difference_hpi.Subtract(&n1, &n2)

		if result.NumberRepresentation.GetFirstNode().GetData().(int) != -1 {
			t.Error("Subtraction has not well performed")
		}
	})

	t.Run("subtract a HPI from another, both with only one node each, subtrahend is greater", func(t *testing.T) {
		n1 := high_precision_integers.SetUpHighPrecisionInteger(1900000000)
		n2 := high_precision_integers.SetUpHighPrecisionInteger(1000065466)
		result := difference_hpi.Subtract(&n1, &n2)

		if result.NumberRepresentation.Size() != 3 {
			t.Error("Subtraction has not well performed")
		}
	})

	t.Run("subtract a HPI from another, minuend much higher than subtrahend", func(t *testing.T) {
		n1 := high_precision_integers.SetUpHighPrecisionInteger(1900000000)
		n2 := high_precision_integers.SetUpHighPrecisionInteger(65466)
		result := difference_hpi.Subtract(&n1, &n2)

		if result.NumberRepresentation.Size() != 3 {
			t.Error("Subtraction has not well performed")
		}
	})

	t.Run("subtract a HPI from another, where firsr number is smaller than the second", func(t *testing.T) {
		n2 := high_precision_integers.SetUpHighPrecisionInteger(1900000000)
		n1 := high_precision_integers.SetUpHighPrecisionInteger(65466)
		result := difference_hpi.Subtract(&n1, &n2)

		if result.NumberRepresentation.Size() != 3 && result.NumberRepresentation.GetFirstNode().GetData().(int) > 0 {
			t.Error("Subtraction has not well performed")
		}
	})

}
