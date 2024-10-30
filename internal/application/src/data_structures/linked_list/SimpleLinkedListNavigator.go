package linked_list

import "fmt"

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

func (simpleLinkedList *SimpleLinkedList) ReverseList() *SimpleLinkedList {
	simpleLinkedList.SwapExtremeNodes()
	if simpleLinkedList.size > 3 {
		leftNode := simpleLinkedList.firstNode.next
		rightNode := simpleLinkedList.getNodeBefore(simpleLinkedList.lastNode)
		for leftNode != rightNode {
			leftNode, rightNode = simpleLinkedList.SwapNodes(leftNode, rightNode)
			fmt.Println(leftNode)
			fmt.Println(rightNode)
		}
	}
	return simpleLinkedList
}

func (simpleLinkedList *SimpleLinkedList) SwapNodes(leftNode *SimpleNode, rightNode *SimpleNode) (*SimpleNode, *SimpleNode) {
	previousLeft := simpleLinkedList.getNodeBefore(leftNode)
	nextToLeft := leftNode.next
	previousRight := simpleLinkedList.getNodeBefore(rightNode)
	nextToRight := rightNode.next

	previousLeft.next = rightNode
	leftNode.next = nextToRight

	previousRight.next = leftNode
	rightNode.next = nextToLeft
	return nextToLeft, previousRight
}

func (simpleLinkedList *SimpleLinkedList) SwapExtremeNodes() {
	if simpleLinkedList.size == 2 {
		temporaryFirstNode := simpleLinkedList.GetFirstNode()
		temporaryLastNode := simpleLinkedList.GetLastNode()
		simpleLinkedList.lastNode = temporaryFirstNode
		simpleLinkedList.firstNode = temporaryLastNode
		temporaryLastNode.next = temporaryFirstNode
		temporaryFirstNode.next = nil
	} else if simpleLinkedList.size > 2 {
		temporaryFirstNode := simpleLinkedList.GetFirstNode()
		nextToFirstNode := temporaryFirstNode.next

		temporaryLastNode := simpleLinkedList.GetLastNode()
		previousOfLastNode := simpleLinkedList.getNodeBefore(temporaryLastNode)

		temporaryLastNode.next = nextToFirstNode
		previousOfLastNode.next = temporaryFirstNode

		simpleLinkedList.firstNode = temporaryLastNode
		simpleLinkedList.lastNode = temporaryFirstNode
		temporaryFirstNode.next = nil
	}
}

func (sll *SimpleLinkedList) getNodeBefore(node *SimpleNode) *SimpleNode {
	nodeRunner := sll.GetFirstNode()
	for nodeRunner != nil {
		if node == nodeRunner.next {
			return nodeRunner
		}
		nodeRunner = nodeRunner.next
	}
	return nil
}
