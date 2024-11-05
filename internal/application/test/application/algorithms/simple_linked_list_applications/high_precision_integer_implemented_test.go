package simple_linked_list_applications

import (
	"github.com/web3foru/computer-science-group-golang/internal/application/src/algorithms/simple_linked_list_applications/high_precision_integers"
	"testing"
)

func TestHighPrecisionIntegerOperations(t *testing.T) {
	t.Run("Reverse a list optimized", func(t *testing.T) {
		newLinkedLinkedList := GetAListForReversing()
		newLinkedLinkedList.ReverseOptimized()
		if newLinkedLinkedList.GetFirstNode().GetData() != 4 {
			t.Error("list hasn't reversed yet")
		}
	})

	t.Run("Represent a high precision integer with the minimum set of numbers", func(t *testing.T) {
		aBigNumber := 9999
		newHPI := high_precision_integers.SetUpHighPrecisionInteger(aBigNumber)
		if newHPI.GetLinkedList().Size() == 0 {
			t.Error("not represented yet")
		}
	})

	t.Run("Represent a high precision integer with two nodes", func(t *testing.T) {
		aBigNumber := 10000
		newHPI := high_precision_integers.SetUpHighPrecisionInteger(aBigNumber)
		if newHPI.GetLinkedList().Size() < 2 {
			t.Error("not represented yet")
		}
	})

	t.Run("Represent a high precision integer with three nodes", func(t *testing.T) {
		aBigNumber := 900000000
		newHPI := high_precision_integers.SetUpHighPrecisionInteger(aBigNumber)
		if newHPI.GetLinkedList().Size() < 3 {
			t.Error("not represented yet")
		}
	})

	t.Run("Represent a high precision integer with random numbers", func(t *testing.T) {
		aBigNumber := 456812359987546345
		newHPI := high_precision_integers.SetUpHighPrecisionInteger(aBigNumber)
		if newHPI.GetLinkedList().Size() < 5 {
			t.Error("not represented yet")
		}
	})
}
