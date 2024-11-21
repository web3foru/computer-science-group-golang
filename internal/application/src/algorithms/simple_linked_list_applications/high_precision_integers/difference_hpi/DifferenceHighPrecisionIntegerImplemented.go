package difference_hpi

import (
	"github.com/web3foru/computer-science-group-golang/internal/application/src/algorithms/simple_linked_list_applications/high_precision_integers"
	"github.com/web3foru/computer-science-group-golang/internal/application/src/data_structures/linked_list"
)

type HPI = high_precision_integers.HighPrecisionInteger

func setUpSubtractionStructure(n1 *HPI, n2 *HPI) *SubtractionStructure {
	n1.GetLinkedList().ReverseOptimized()
	n2.GetLinkedList().ReverseOptimized()
	minuend := n1
	subtrahend := n2
	if n1.NumberRepresentation.Size() < n2.NumberRepresentation.Size() {
		minuend = n2
		subtrahend = n1
	}
	result := linked_list.NewLinkedList()
	return &SubtractionStructure{minuend, subtrahend, n1.GetLinkedList().GetFirstNode(), n2.GetLinkedList().GetFirstNode(), 0, result}
}

func Subtract(n1 *HPI, n2 *HPI) *HPI {
	subtractionStructure := setUpSubtractionStructure(n1, n2)
	for thereAreElementsInBothIntegers(subtractionStructure) {
		subtractUntilOneOfThemComeToTheLast(subtractionStructure)
	}

	return &HPI{subtractionStructure.result}
}

func finishSubtractionByMinuend(subtractionStructure *SubtractionStructure) {
	for subtractionStructure.minuendPointer != nil {

	}
}

func thereAreElementsInBothIntegers(subtractionStructure *SubtractionStructure) bool {
	return subtractionStructure.minuendPointer != nil && subtractionStructure.subtrahendPointer != nil
}

func subtractUntilOneOfThemComeToTheLast(subtractionStructure *SubtractionStructure) {
	minuendPointer := subtractionStructure.minuendPointer
	subtrahendPointer := subtractionStructure.subtrahendPointer
	borrow := subtractionStructure.borrow
	temporaryMinuendValue := minuendPointer.GetData().(int)
	if borrow > 0 {
		temporaryMinuendValue -= 1
		if temporaryMinuendValue < 0 {
			borrow = 1
			temporaryMinuendValue = 9999
		} else {
			borrow = 0
		}
	}
	resultOfOperation := temporaryMinuendValue - subtrahendPointer.GetData().(int)

	if resultOfOperation < 0 {
		resultOfOperation = resultOfOperation + 10000
		borrow = 1
	}
	subtractionStructure.result.AddNodeAtTheBeginning(resultOfOperation)
	subtractionStructure.borrow = borrow
	updatePointers(subtractionStructure)
}

func updatePointers(subtractionStructure *SubtractionStructure) {
	subtractionStructure.minuendPointer = subtractionStructure.minuendPointer.GetNextNode()
	subtractionStructure.subtrahendPointer = subtractionStructure.subtrahendPointer.GetNextNode()

}
