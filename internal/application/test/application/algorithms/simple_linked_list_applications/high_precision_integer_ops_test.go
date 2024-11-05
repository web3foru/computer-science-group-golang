package simple_linked_list_applications

import (
	"github.com/web3foru/computer-science-group-golang/internal/application/src/algorithms/simple_linked_list_applications/high_precision_integers"
	"github.com/web3foru/computer-science-group-golang/internal/application/src/algorithms/simple_linked_list_applications/high_precision_integers/sum_hpi"
	"testing"
)

func TestSumOperations(t *testing.T) {
	t.Run("Sum two hpis with one single node each one", func(t *testing.T) {
		numberOne := high_precision_integers.SetUpHighPrecisionInteger(100)
		numberTwo := high_precision_integers.SetUpHighPrecisionInteger(100)
		result := sum_hpi.Sum(&numberTwo, &numberOne)
		if result.GetLinkedList().GetFirstNode().GetData() != 200 {
			t.Error("Operation wasn't performed successfully")
		}
	})

	t.Run("Sum two hpis with two nodes", func(t *testing.T) {
		numberOne := high_precision_integers.SetUpHighPrecisionInteger(19000)
		numberTwo := high_precision_integers.SetUpHighPrecisionInteger(10000)
		result := sum_hpi.Sum(&numberTwo, &numberOne)
		if result.GetLinkedList().Size() != 2 {
			t.Error("Operation wasn't performed successfully")
		}
	})

	t.Run("Sum two hpis with a large number", func(t *testing.T) {
		numberOne := high_precision_integers.SetUpHighPrecisionInteger(19000000)
		numberTwo := high_precision_integers.SetUpHighPrecisionInteger(1000065466)
		result := sum_hpi.Sum(&numberTwo, &numberOne)
		if result.GetLinkedList().Size() != 3 {
			t.Error("Operation wasn't performed successfully")
		}
	})
}
