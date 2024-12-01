package hpi_multiplication

import (
	"github.com/web3foru/computer-science-group-golang/internal/application/src/algorithms/simple_linked_list_applications/high_precision_integers"
	"github.com/web3foru/computer-science-group-golang/internal/application/src/algorithms/simple_linked_list_applications/high_precision_integers/multiplication_hpi"
	"testing"
)

func TestMultiplicationHPI(t *testing.T) {
	t.Run("basic multiplication", func(t *testing.T) {
		n1 := high_precision_integers.SetUpHighPrecisionInteger(999)
		n2 := high_precision_integers.SetUpHighPrecisionInteger(10)
		result := multiplication_hpi.Multiply(&n1, &n2)
		if result.NumberRepresentation.Size() != 1 {
			t.Error("multiplication has not implemented yer")
		}
	})
}
