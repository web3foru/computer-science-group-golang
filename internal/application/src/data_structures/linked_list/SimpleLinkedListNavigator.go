package linked_list

func (simpleLinkedList *SimpleLinkedList) isTheEndOfTheList(node *SimpleNode) bool {
	if node == nil {
		return true
	} else {
		return false
	}
}

func (simpleLinkedList *SimpleLinkedList) goToTheEndOfList() *SimpleNode {
	currentNode := simpleLinkedList.firstNode
	for !simpleLinkedList.isTheEndOfTheList(currentNode) {
		if currentNode.next == nil {
			return currentNode
		}
		currentNode = currentNode.next
	}
	return currentNode
}
