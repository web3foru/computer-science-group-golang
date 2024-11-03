package simple_linked_list_applications

import (
	"github.com/web3foru/computer-science-group-golang/internal/application/src/algorithms/simple_linked_list_applications"
	"testing"
)

func TestSumOperations(t *testing.T) {
	t.Run("Sum two hpis with one single node each one", func(t *testing.T) {
		numberOne := simple_linked_list_applications.SetUpHighPrecisionInteger(100)
		numberTwo := simple_linked_list_applications.SetUpHighPrecisionInteger(100)
		result := numberTwo.Sum(numberOne)
		if result.GetLinkedList().GetFirstNode().GetData() != 200 {
			t.Error("Operation wasn't performed successfully")
		}
	})

	t.Run("Sum two hpis with two nodes", func(t *testing.T) {
		numberOne := simple_linked_list_applications.SetUpHighPrecisionInteger(19000)
		numberTwo := simple_linked_list_applications.SetUpHighPrecisionInteger(10000)
		result := numberTwo.Sum(numberOne)
		if result.GetLinkedList().Size() != 2 {
			t.Error("Operation wasn't performed successfully")
		}
	})

	t.Run("Sum two hpis with a large number", func(t *testing.T) {
		numberOne := simple_linked_list_applications.SetUpHighPrecisionInteger(19000000)
		numberTwo := simple_linked_list_applications.SetUpHighPrecisionInteger(1000065466)
		result := numberTwo.Sum(numberOne)
		if result.GetLinkedList().Size() != 2 {
			t.Error("Operation wasn't performed successfully")
		}
	})
}
