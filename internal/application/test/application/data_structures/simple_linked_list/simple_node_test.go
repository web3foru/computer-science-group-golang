package simple_linked_list

import (
	"github.com/web3foru/computer-science-group-golang/internal/application/src/data_structures/linked_list"
	"testing"
)

func TestSimpleNode(t *testing.T) {
	t.Run("Testing a new empty node", func(t *testing.T) {
		newSimpleNode := linked_list.CreateNewNode("hello!")
		dataOfNewNode := newSimpleNode.String()
		if dataOfNewNode != "hello!" {
			t.Error("The Node did not set up correctly")
		}
	})
}
