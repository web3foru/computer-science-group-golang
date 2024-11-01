package linked_list

func (simpleLinkedList *SimpleLinkedList) disconnectNodes(currentNode *SimpleNode, previousNode *SimpleNode) bool {
	previousNode.next = currentNode.next
	currentNode = nil
	simpleLinkedList.size--
	return true
}

func (simpleLinkedList *SimpleLinkedList) removeNodeAtTheBeginning() bool {
	simpleLinkedList.firstNode = simpleLinkedList.firstNode.next
	return true
}

func (simpleLinkedList *SimpleLinkedList) RemoveNode(data interface{}) bool {
	previousNode := simpleLinkedList.firstNode
	currentNode := previousNode.next
	for !simpleLinkedList.IsTheEndOfTheList(currentNode) {
		if currentNode.data == data {
			return simpleLinkedList.disconnectNodes(currentNode, previousNode)
		} else {
			previousNode = previousNode.next
			currentNode = currentNode.next
		}
	}
	return false
}

func (simpleLinkedList *SimpleLinkedList) ReturnFirstNodeAndErase() *SimpleNode {
	currentNode := simpleLinkedList.firstNode
	simpleLinkedList.firstNode = currentNode.next
	simpleLinkedList.size--
	return currentNode
}

func (simpleLinkedList *SimpleLinkedList) ReturnLastNodeAndErase() *SimpleNode {
	if simpleLinkedList.IsEmpty() {
		return nil
	} else {
		previous, lastNode := simpleLinkedList.GetTheLastTwoNodes()
		if lastNode != nil {
			previous.next = nil
			simpleLinkedList.size--
		}
		return lastNode
	}

}
