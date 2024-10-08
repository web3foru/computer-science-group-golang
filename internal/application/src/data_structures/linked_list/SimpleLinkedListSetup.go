package linked_list

func NewLinkedList() *SimpleLinkedList {
	return &SimpleLinkedList{firstNode: nil, size: 0}
}

func (simpleLinkedList *SimpleLinkedList) GetFirstNode() *SimpleNode {
	return simpleLinkedList.firstNode
}

func (simpleLinkedList *SimpleLinkedList) IsEmpty() bool {
	return simpleLinkedList.firstNode == nil
}

func (simpleLinkedList *SimpleLinkedList) Size() int {
	return simpleLinkedList.size
}
