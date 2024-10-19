package linked_list

func (list *CircularLinkedList) AddNodeAtBeginning(data interface{}) {
	newNode := &CircularNode{data: data}
	if list.IsEmpty() {
		list.head.next = newNode
		newNode.next = list.head
	} else {
		newNode.next = list.head.next
		list.head.next = newNode
	}
	list.size++
}

func (list *CircularLinkedList) AddNodeAtEnd(data interface{}) {
	newNode := &CircularNode{data: data}
	if list.IsEmpty() {
		list.head.next = newNode
		newNode.next = list.head
	} else {
		tail := list.getTailNode()
		tail.next = newNode
		newNode.next = list.head
	}
	list.size++
}

func (list *CircularLinkedList) getTailNode() *CircularNode {
	current := list.head.next
	for current.next != list.head {
		current = current.next
	}
	return current
}
