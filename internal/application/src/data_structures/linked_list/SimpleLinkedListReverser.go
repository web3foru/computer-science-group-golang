package linked_list

func (sll *SimpleLinkedList) ReverseOptimized() {
	prev := CreateNewNode(nil)
	current := sll.GetFirstNode()
	sll.lastNode = current
	for current != nil {
		nextNode := current.GetNextNode()
		current.next = prev
		prev = current
		current = nextNode
	}
	sll.lastNode.next = nil
	sll.firstNode = prev
}
