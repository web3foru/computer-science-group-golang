package simple_linked_list

import (
	"testing"
)

func TestLinkedListOperations(t *testing.T) {

	t.Run("Insert a new element at beginning", func(t *testing.T) {
		newLinked := GetAnSimpleListWithOneNode()
		if newLinked.IsEmpty() {
			t.Errorf("New linked list should not be empty")
		}
	})

	t.Run("Insert a new element at beginning of an existing list", func(t *testing.T) {
		newLinked := GetAnSimpleListWithOneNode()
		newLinked.AddNodeAtTheBeginning("a new node x2")
		if newLinked.Size() != 2 {
			t.Errorf("linked list must has two elements")
		}
	})

	t.Run("Insert a new element at the end of an existing list", func(t *testing.T) {
		newLinked := GetAnSimpleListWithOneNode()
		newLinked.AddNodeAtTheEnd("a new node x2")
		firstNode := newLinked.GetFirstNode()

		if firstNode.String() != "a new node" {
			t.Errorf("the first node must match with the exactly string")
		}

		nextNode := firstNode.GetNextNode()
		if nextNode.String() != "a new node x2" {
			t.Errorf("the next node must match with the exactly string")
		}
	})

	t.Run("Delete one node based on data", func(t *testing.T) {
		newLinked := GetAnSimpleListWithNNodes(10)

		newLinked.RemoveNode(5)
		if newLinked.Size() != 9 {
			t.Errorf("linked list must has 9 elements but has %v", newLinked.Size())
		}
	})
}
