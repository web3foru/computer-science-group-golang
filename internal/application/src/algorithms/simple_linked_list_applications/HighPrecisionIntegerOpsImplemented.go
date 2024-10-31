package simple_linked_list_applications

import "github.com/web3foru/computer-science-group-golang/internal/application/src/data_structures/linked_list"

func (number1 *HighPrecisionInteger) Sum(number2 HighPrecisionInteger) HighPrecisionInteger {
	result := linked_list.NewLinkedList()
	carryOut := linked_list.CreateNewNode(0)
	dataForInsertion := 0
	firstNumberPointerNode := number1.getLinkedList().GetFirstNode()
	secondNumberPointerNode := number2.getLinkedList().GetFirstNode()
	for !number1.numberRepresentation.IsTheEndOfTheList(firstNumberPointerNode) &&
		!number2.numberRepresentation.IsTheEndOfTheList(secondNumberPointerNode) {
		carryOut, dataForInsertion = evaluateSum(firstNumberPointerNode, secondNumberPointerNode, carryOut)
		firstNumberPointerNode = firstNumberPointerNode.GetNextNode()
		secondNumberPointerNode = secondNumberPointerNode.GetNextNode()
		result.AddNodeAtTheBeginning(dataForInsertion)
	}

	return HighPrecisionInteger{numberRepresentation: result}
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
