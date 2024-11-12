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
}
