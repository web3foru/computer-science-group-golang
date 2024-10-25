package linked_list

func (list *DoublyLinkedList) RemoveNode(data interface{}) bool {
	current := list.head
	for current != nil {
		if current.data == data {
			if current == list.head {
				list.RemoveNodeAtBeginning()
			} else if current == list.tail {
				list.RemoveNodeAtEnd()
			} else {
				current.prev.next = current.next
				current.next.prev = current.prev
				list.size--
			}
			return true
		}
		current = current.next
	}
	return false
}

func (list *DoublyLinkedList) RemoveNodeAtBeginning() bool {
	if list.IsEmpty() {
		return false
	}
	if list.head == list.tail {
		list.head = nil
		list.tail = nil
	} else {
		list.head = list.head.next
		list.head.prev = nil
	}
	list.size--
	return true
}

func (list *DoublyLinkedList) RemoveNodeAtEnd() bool {
	if list.IsEmpty() {
		return false
	}
	if list.head == list.tail {
		list.head = nil
		list.tail = nil
	} else {
		list.tail = list.tail.prev
		list.tail.next = nil
	}
	list.size--
	return true
}
