package linked_list

func (list *DoublyLinkedList) IsEnd(node *DoubleNode) bool {
	return node == nil
}

func (list *DoublyLinkedList) IterateForward(f func(data interface{})) {
	for current := list.head; current != nil; current = current.next {
		f(current.data)
	}
}

func (list *DoublyLinkedList) IterateBackward(f func(data interface{})) {
	for current := list.tail; current != nil; current = current.prev {
		f(current.data)
	}
}
