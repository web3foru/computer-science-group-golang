package linked_list

func (simpleLinkedList *SimpleLinkedList) AddNodeAtTheBeginning(data interface{}) bool {
	newNode := CreateNewNode(data)
	simpleLinkedList.updateElementsAtBeginning(newNode)
	return true
}

func (simpleLinkedList *SimpleLinkedList) updateElementsAtBeginning(newNode *SimpleNode) {
	newNode.next = simpleLinkedList.firstNode
	simpleLinkedList.firstNode = newNode
	simpleLinkedList.size++
}

func (simpleLinkedList *SimpleLinkedList) AddNodeAtTheEnd(data interface{}) {
	newNode := CreateNewNode(data)
	simpleLinkedList.updateElementsAtTheEnd(newNode)
	simpleLinkedList.size++
}

func (simpleLinkedList *SimpleLinkedList) updateElementsAtTheEnd(newNode *SimpleNode) {
	lastNode := simpleLinkedList.goToTheEndOfList()
	if lastNode != nil {
		lastNode.next = newNode
	} else {
		simpleLinkedList.firstNode = newNode
	}
}
