package linked_list

func (list *CircularLinkedList) Iterate(f func(data interface{})) {
	if list.IsEmpty() {
		return
	}
	current := list.head.next
	for current != list.head {
		f(current.data)
		current = current.next
	}
}
