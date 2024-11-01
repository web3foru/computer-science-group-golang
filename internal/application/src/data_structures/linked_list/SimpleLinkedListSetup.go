package linked_list

func NewLinkedList() *SimpleLinkedList {
	return &SimpleLinkedList{firstNode: nil, lastNode: nil, size: 0}
}

func (simpleLinkedList *SimpleLinkedList) GetFirstNode() *SimpleNode {
	return simpleLinkedList.firstNode
}

func (simpleLinkedList *SimpleLinkedList) GetLastNode() *SimpleNode {
	return simpleLinkedList.lastNode
}

func (simpleLinkedList *SimpleLinkedList) IsEmpty() bool {
	return simpleLinkedList.firstNode == nil
}

func (simpleLinkedList *SimpleLinkedList) Size() int {
	return simpleLinkedList.size
}

func (simpleLinkedList *SimpleLinkedList) SetFirstNode(newFirstNode *SimpleNode) {
	if simpleLinkedList.firstNode == nil {
		simpleLinkedList.AddNodeAtTheBeginning(newFirstNode)
	} else {
		previousFirstNode := simpleLinkedList.GetFirstNode()
		newFirstNode.UpdateNext(previousFirstNode)
		simpleLinkedList.firstNode = newFirstNode
	}
}
