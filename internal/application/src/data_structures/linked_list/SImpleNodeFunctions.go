package linked_list

import (
	"fmt"
)

func (simpleNode *SimpleNode) String() string {
	return fmt.Sprintf("%v", simpleNode.data)
}

func (simpleNode *SimpleNode) GetData() interface{} {
	return simpleNode.data
}

func (simpleNode *SimpleNode) GetNextNode() *SimpleNode {
	return simpleNode.next
}

func (simpleNode *SimpleNode) UpdateNext(nextNode *SimpleNode) {
	simpleNode.next = nextNode
}

func CreateNewNode(data interface{}) *SimpleNode {
	return &SimpleNode{data: data, next: nil}
}
