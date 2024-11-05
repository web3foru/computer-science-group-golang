package sum_hpi

import (
	"github.com/web3foru/computer-science-group-golang/internal/application/src/algorithms/simple_linked_list_applications/high_precision_integers"
	"github.com/web3foru/computer-science-group-golang/internal/application/src/data_structures/linked_list"
)

func Sum(number1 *high_precision_integers.HighPrecisionInteger, number2 *high_precision_integers.HighPrecisionInteger) high_precision_integers.HighPrecisionInteger {
	result := linked_list.NewLinkedList()
	carryOut := linked_list.CreateNewNode(0)
	dataForInsertion := 0
	number1.GetLinkedList().ReverseOptimized()
	number2.GetLinkedList().ReverseOptimized()
	firstNumberPointerNode := number1.GetLinkedList().GetFirstNode()
	secondNumberPointerNode := number2.GetLinkedList().GetFirstNode()

	for !number1.GetLinkedList().IsTheEndOfTheList(firstNumberPointerNode) &&
		!number2.GetLinkedList().IsTheEndOfTheList(secondNumberPointerNode) {
		carryOut, dataForInsertion = evaluateSum(firstNumberPointerNode, secondNumberPointerNode, carryOut)
		firstNumberPointerNode = firstNumberPointerNode.GetNextNode()
		secondNumberPointerNode = secondNumberPointerNode.GetNextNode()
		result.AddNodeAtTheBeginning(dataForInsertion)
	}
	evaluateRemainPartOfTheNumber(number1, firstNumberPointerNode, carryOut, result)
	evaluateRemainPartOfTheNumber(number2, secondNumberPointerNode, carryOut, result)
	addCarryOut(result, carryOut)
	return high_precision_integers.HighPrecisionInteger{NumberRepresentation: result}
}

func addCarryOut(result *linked_list.SimpleLinkedList, carryOut *linked_list.SimpleNode) {
	if carryOut.GetData() != 0 {
		result.AddNodeAtTheBeginning(carryOut.GetData())
	}
}

func evaluateRemainPartOfTheNumber(number *high_precision_integers.HighPrecisionInteger, checkpointNode *linked_list.SimpleNode,
	carryOut *linked_list.SimpleNode, currentResult *linked_list.SimpleLinkedList) *linked_list.SimpleNode {
	dataForInsertion := 0
	for !number.NumberRepresentation.IsTheEndOfTheList(checkpointNode) {
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
