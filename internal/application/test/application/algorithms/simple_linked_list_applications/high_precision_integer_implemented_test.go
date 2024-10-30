package simple_linked_list_applications

import (
	"testing"
)

func TestHighPrecisionIntegerOperations(t *testing.T) {
	t.Run("should reverse a simple linked list", func(t *testing.T) {
		newLinkedLinkedList := GetAListForReversing()
		newLinkedLinkedList.ReverseList()
		if newLinkedLinkedList.GetFirstNode().GetData() != 4 {
			t.Error("list hasn't reversed yet")
		}
	})

	t.Run("should swap the first and the last node", func(t *testing.T) {
		newLinkedLinkedList := GetAListWithTwoNodesForSwaping()
		newLinkedLinkedList.SwapExtremeNodes()
		if newLinkedLinkedList.GetFirstNode().GetData() != 1 {
			t.Error("list hasn't swaped yet")
		}
	})

	t.Run("should swap the first and the last node of a longest List", func(t *testing.T) {
		newLinkedLinkedList := GetAListForReversing()
		newLinkedLinkedList.SwapExtremeNodes()
		if newLinkedLinkedList.GetFirstNode().GetData() != 4 {
			t.Error("list hasn't swaped yet")
		}
	})
}
