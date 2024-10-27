package linked_list

func (simpleLinkedList *SimpleLinkedList) IsTheEndOfTheList(node *SimpleNode) bool {
	if node == nil {
		return true
	} else {
		return false
	}
}

func (simpleLinkedList *SimpleLinkedList) GoToTheEndOfList() *SimpleNode {
	currentNode := simpleLinkedList.firstNode
	for !simpleLinkedList.IsTheEndOfTheList(currentNode) {
		if currentNode.next == nil {
			return currentNode
		}
		currentNode = currentNode.next
	}
	return currentNode
}

func (simpleLinkedList *SimpleLinkedList) GetTheLastTwoNodes() (*SimpleNode, *SimpleNode) {
	currentNode := simpleLinkedList.firstNode
	if currentNode == nil {
		return nil, nil
	} else if currentNode.next == nil {
		return currentNode, nil
	} else {
		nextNode := currentNode.next
		for !simpleLinkedList.IsTheEndOfTheList(nextNode) {
			if nextNode.next == nil {
				return currentNode, nextNode
			} else {
				currentNode = currentNode.next
				nextNode = nextNode.next
			}
		}
	}
	return nil, nil
}
