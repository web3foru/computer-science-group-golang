package simple_linked_list_applications

import "github.com/web3foru/computer-science-group-golang/internal/application/src/data_structures/linked_list"

func (number1 *HighPrecisionInteger) Sum(number2 HighPrecisionInteger) HighPrecisionInteger {
	result := linked_list.NewLinkedList()
	carryOut := linked_list.CreateNewNode(0)
	dataForInsertion := 0
	number1.GetLinkedList().ReverseOptimized()
	number2.GetLinkedList().ReverseOptimized()
	firstNumberPointerNode := number1.getLinkedList().GetFirstNode()
	secondNumberPointerNode := number2.getLinkedList().GetFirstNode()

	for !number1.numberRepresentation.IsTheEndOfTheList(firstNumberPointerNode) &&
		!number2.numberRepresentation.IsTheEndOfTheList(secondNumberPointerNode) {
		carryOut, dataForInsertion = evaluateSum(firstNumberPointerNode, secondNumberPointerNode, carryOut)
		firstNumberPointerNode = firstNumberPointerNode.GetNextNode()
		secondNumberPointerNode = secondNumberPointerNode.GetNextNode()
		result.AddNodeAtTheBeginning(dataForInsertion)
	}

	for !number1.numberRepresentation.IsTheEndOfTheList(firstNumberPointerNode) {
		carryOut, dataForInsertion = evaluateSum(firstNumberPointerNode, linked_list.CreateNewNode(0), carryOut)
		result.AddNodeAtTheBeginning(dataForInsertion)
		firstNumberPointerNode = firstNumberPointerNode.GetNextNode()
	}

	for !number2.numberRepresentation.IsTheEndOfTheList(secondNumberPointerNode) {
		carryOut, dataForInsertion = evaluateSum(linked_list.CreateNewNode(0), secondNumberPointerNode, carryOut)
		result.AddNodeAtTheBeginning(dataForInsertion)
		secondNumberPointerNode = secondNumberPointerNode.GetNextNode()
	}
	if carryOut.GetData() != 0 {
		result.AddNodeAtTheBeginning(carryOut.GetData())
	}
	return HighPrecisionInteger{numberRepresentation: result}
}

func evaluateRemainPartOfTheFirstNumber(number1 *HighPrecisionInteger, checkpointNode *linked_list.SimpleNode, carryOut *linked_list.SimpleNode, currentResult *linked_list.SimpleLinkedList) *linked_list.SimpleNode {
	dataForInsertion := 0
	for !number1.numberRepresentation.IsTheEndOfTheList(checkpointNode) {
		carryOut, dataForInsertion = evaluateSum(checkpointNode, linked_list.CreateNewNode(0), carryOut)
		currentResult.AddNodeAtTheBeginning(dataForInsertion)
		checkpointNode = checkpointNode.GetNextNode()
	}
	return carryOut
}

func evaluateSum(node1 *linked_list.SimpleNode, node2 *linked_list.SimpleNode, carryOut *linked_list.SimpleNode) (*linked_list.SimpleNode, int) {
	temporaryResult := node1.GetData().(int) + node2.GetData().(int) + carryOut.GetData().(int)
	carryOut.UpdateData(temporaryResult / 10000)
	dataForInsertion := temporaryResult % 10000
	return carryOut, dataForInsertion
}

func (hpi *HighPrecisionInteger) Difference(number2 HighPrecisionInteger) HighPrecisionInteger {
	return SetUpHighPrecisionInteger(1)
}
