package linked_list

func (list *DoublyLinkedList) AddNodeAtBeginning(data interface{}) {
	newNode := CreateNewNode(data)
	if list.IsEmpty() {
		list.head = newNode
		list.tail = newNode
	} else {
		newNode.next = list.head
		list.head.prev = newNode
		list.head = newNode
	}
	list.size++
}

func (list *DoublyLinkedList) AddNodeAtEnd(data interface{}) {
	newNode := CreateNewNode(data)
	if list.IsEmpty() {
		list.head = newNode
		list.tail = newNode
	} else {
		newNode.prev = list.tail
		list.tail.next = newNode
		list.tail = newNode
	}
	list.size++
}
