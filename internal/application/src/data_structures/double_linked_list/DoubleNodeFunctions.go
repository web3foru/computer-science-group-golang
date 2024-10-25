package linked_list

import "fmt"

func (node *DoubleNode) String() string {
	return fmt.Sprintf("%v", node.data)
}

func (node *DoubleNode) GetNextNode() *DoubleNode {
	return node.next
}

func (node *DoubleNode) GetPrevNode() *DoubleNode {
	return node.prev
}

func CreateNewNode(data interface{}) *DoubleNode {
	return &DoubleNode{data: data, next: nil, prev: nil}
}
