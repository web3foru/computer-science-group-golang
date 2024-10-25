package linked_list

import "fmt"

func (node *CircularNode) String() string {
	return fmt.Sprintf("%v", node.data)
}

func (node *CircularNode) GetNextNode() *CircularNode {
	return node.next
}
