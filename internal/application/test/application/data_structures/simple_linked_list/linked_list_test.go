package simple_linked_list

import (
	"testing"
)

func TestLinkedList(t *testing.T) {

	t.Run("An Empty list must be created", func(t *testing.T) {
		oneMockList := GetAnEmptySimpleList()
		if !(oneMockList.IsEmpty() || oneMockList.Size() == 0) {
			t.Errorf("Expected an empty list")
		}
	})
}
