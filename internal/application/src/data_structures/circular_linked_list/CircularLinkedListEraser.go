package linked_list

func (list *CircularLinkedList) RemoveNode(data interface{}) bool {
	if list.IsEmpty() {
		return false
	}

	previous := list.head
	current := list.head.next

	for current != list.head {
		if current.data == data {
			previous.next = current.next
			list.size--
			return true
		}
		previous = current
		current = current.next
	}
	return false
}

func (list *CircularLinkedList) RemoveNodeAtBeginning() bool {
	if list.IsEmpty() {
		return false
	}
	firstNode := list.head.next
	list.head.next = firstNode.next
	list.size--
	return true
}

func (list *CircularLinkedList) RemoveNodeAtEnd() bool {
	if list.IsEmpty() {
		return false
	}
	previous := list.head
	current := list.head.next

	for current.next != list.head {
		previous = current
		current = current.next
	}
	previous.next = list.head
	list.size--
	return true
}
