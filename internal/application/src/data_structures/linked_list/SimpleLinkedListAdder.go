package linked_list

func (simpleLinkedList *SimpleLinkedList) AddNodeAtTheBeginning(data interface{}) bool {
	newNode := CreateNewNode(data)
	simpleLinkedList.updateElementsAtBeginning(newNode)
	simpleLinkedList.size++
	return true
}

func (simpleLinkedList *SimpleLinkedList) updateElementsAtBeginning(newNode *SimpleNode) {
	if simpleLinkedList.IsEmpty() {
		simpleLinkedList.firstNode = newNode
		simpleLinkedList.lastNode = newNode
	}
	newNode.next = simpleLinkedList.firstNode
	simpleLinkedList.firstNode = newNode
}

func (simpleLinkedList *SimpleLinkedList) AddNodeAtTheEnd(data interface{}) {
	newNode := CreateNewNode(data)
	if simpleLinkedList.IsEmpty() {
		simpleLinkedList.size = 1
	} else {
		simpleLinkedList.size++
	}
	simpleLinkedList.ReplaceElementAtTheEnd(newNode)
}

func (simpleLinkedList *SimpleLinkedList) setUpEmptyList(data interface{}) {

}

func (simpleLinkedList *SimpleLinkedList) InsertAfter(data interface{}, newNodeData interface{}) bool {

	currentNode := simpleLinkedList.firstNode
	for !simpleLinkedList.IsTheEndOfTheList(currentNode) {
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

func (simpleLinkedList *SimpleLinkedList) ReplaceElementAtTheEnd(newNode *SimpleNode) {
	if simpleLinkedList.IsEmpty() {
		simpleLinkedList.SetupEmptyList(newNode)
	} else if newNode == simpleLinkedList.firstNode {
		newFirstNode := simpleLinkedList.firstNode.next
		simpleLinkedList.firstNode = newFirstNode
		previousLastNode := simpleLinkedList.lastNode
		previousLastNode.next = newNode
		simpleLinkedList.lastNode = newNode
		newNode.next = nil
	} else {
		previousLastNode := simpleLinkedList.GetLastNode()
		simpleLinkedList.lastNode = newNode
		previousLastNode.next = newNode
	}

}

func (simpleLinkedList *SimpleLinkedList) ReplaceElementAtBeginning(newNode *SimpleNode) {
	if simpleLinkedList.IsEmpty() {
		simpleLinkedList.SetupEmptyList(newNode)
	} else {
		previousFirstNode := simpleLinkedList.firstNode
		simpleLinkedList.firstNode = newNode
		newNode.next = previousFirstNode
	}
}

func (simpleLinkedList *SimpleLinkedList) SetupEmptyList(newNode *SimpleNode) {
	simpleLinkedList.firstNode = newNode
	simpleLinkedList.lastNode = newNode
	simpleLinkedList.size = 1
}
