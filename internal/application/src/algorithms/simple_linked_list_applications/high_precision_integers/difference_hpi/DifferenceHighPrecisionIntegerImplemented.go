package difference_hpi

import (
	"github.com/web3foru/computer-science-group-golang/internal/application/src/algorithms/simple_linked_list_applications/high_precision_integers"
	"github.com/web3foru/computer-science-group-golang/internal/application/src/data_structures/linked_list"
)

type HPI = high_precision_integers.HighPrecisionInteger

func setUpSubtractionStructure(n1 *HPI, n2 *HPI) *SubtractionStructure {
	borrow := linked_list.CreateNewNode(0)
	n1.GetLinkedList().ReverseOptimized()
	n2.GetLinkedList().ReverseOptimized()
	result := linked_list.NewLinkedList()
	return &SubtractionStructure{n1, n2, n1.GetLinkedList().GetFirstNode(), n2.GetLinkedList().GetFirstNode(), borrow, result}
}

func Subtract(n1 *HPI, n2 *HPI) *HPI {
	subtractionStructure := setUpSubtractionStructure(n1, n2)
	for thereAreElementsInBothIntegers(subtractionStructure) {
		subtractByChosenNodes(subtractionStructure)
	}
	return &HPI{subtractionStructure.result}
}

func thereAreElementsInBothIntegers(subtractionStructure *SubtractionStructure) bool {
	return subtractionStructure.minuendPointer != nil && subtractionStructure.subtrahendPointer != nil
}

func subtractByChosenNodes(subtractionStructure *SubtractionStructure) {
	result := subtractionStructure.minuendPointer.GetData().(int) - subtractionStructure.subtrahendPointer.GetData().(int)
	subtractionStructure.result.AddNodeAtTheBeginning(result)
	updatePointers(subtractionStructure)
}

func updatePointers(subtractionStructure *SubtractionStructure) {
	subtractionStructure.minuendPointer = subtractionStructure.minuendPointer.GetNextNode()
	subtractionStructure.subtrahendPointer = subtractionStructure.subtrahendPointer.GetNextNode()
}
