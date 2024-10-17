package simple_linked_list

import "testing"

func TestLinkedListAdder(t *testing.T) {
	t.Run("Add element to a list with a single node", func(t *testing.T) {
		singleNodeLinkedList := GetAnSimpleListWithNNodes(1)
		resultOfInsertion := singleNodeLinkedList.InsertAfter(0, 1)
		if !(singleNodeLinkedList.Size() == 2 && resultOfInsertion) {
			t.Error("Expected size of a single node to equal 2")
		}
	})

	t.Run("Add element to a list with two nodes in the middle", func(t *testing.T) {
		singleNodeLinkedList := GetAnSimpleListWithNNodes(2)
		resultOfInsertion := singleNodeLinkedList.InsertAfter(0, 3)
		if !(singleNodeLinkedList.Size() == 3 && resultOfInsertion) {
			t.Error("Expected size of a single node to equal 3")
		}
	})

	t.Run("Add element to a list with two nodes in the end", func(t *testing.T) {
		singleNodeLinkedList := GetAnSimpleListWithNNodes(2)
		resultOfInsertion := singleNodeLinkedList.InsertAfter(1, 2)
		if !(singleNodeLinkedList.Size() == 3 && resultOfInsertion) {
			t.Error("Expected size of a single node to equal 3")
		}
	})

}
