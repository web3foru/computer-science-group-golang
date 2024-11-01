package linked_list

func NewCircularLinkedList() *CircularLinkedList {
	headNode := &CircularNode{data: nil}
	headNode.next = headNode // Apunta a sí mismo inicialmente
	return &CircularLinkedList{head: headNode, size: 0};
}

func (list *CircularLinkedList) IsEmpty() bool {
	return list.size == 0
}

func (list *CircularLinkedList) Size() int {
	return list.size
}
