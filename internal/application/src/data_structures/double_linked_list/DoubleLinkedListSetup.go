package linked_list

func NewDoublyLinkedList() *DoublyLinkedList {
	return &DoublyLinkedList{head: nil, tail: nil, size: 0}
}

func (list *DoublyLinkedList) GetHead() *DoubleNode {
	return list.head
}

func (list *DoublyLinkedList) GetTail() *DoubleNode {
	return list.tail
}

func (list *DoublyLinkedList) IsEmpty() bool {
	return list.head == nil
}

func (list *DoublyLinkedList) Size() int {
	return list.size
}
