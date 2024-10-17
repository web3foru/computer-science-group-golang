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

func (simpleLinkedList *SimpleLinkedList) InsertAfter(data interface{}, newNodeData interface{}) bool {

	currentNode := simpleLinkedList.firstNode
	for !simpleLinkedList.isTheEndOfTheList(currentNode) {
		if currentNode.data == data {
			nextNode := currentNode.next
			newNode := CreateNewNode(newNodeData)
			currentNode.next = newNode
			newNode.next = nextNode
			simpleLinkedList.size++
			return true
		}
		currentNode = currentNode.next
	}
	return false
}

func (simpleLinkedList *SimpleLinkedList) updateElementsAtTheEnd(newNode *SimpleNode) {
	lastNode := simpleLinkedList.goToTheEndOfList()
	if lastNode != nil {
		lastNode.next = newNode
	} else {
		simpleLinkedList.firstNode = newNode
	}
}
