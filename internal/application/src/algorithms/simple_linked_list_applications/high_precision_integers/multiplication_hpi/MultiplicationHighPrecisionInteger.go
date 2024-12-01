package multiplication_hpi

import (
	"github.com/web3foru/computer-science-group-golang/internal/application/src/algorithms/simple_linked_list_applications/high_precision_integers"
	"github.com/web3foru/computer-science-group-golang/internal/application/src/data_structures/linked_list"
)

type HPI = high_precision_integers.HighPrecisionInteger

func Multiply(a, b *HPI) *HPI {
	firstNumberPointer := a.NumberRepresentation.GetFirstNode()
	secondNumberPointer := b.NumberRepresentation.GetFirstNode()
	result := linked_list.NewLinkedList()
	currentOffset := 0
	carryOut := 0
	for firstNumberPointer != nil {
		for secondNumberPointer != nil {
			currentResult := evaluateCurrentMultiplication(firstNumberPointer, secondNumberPointer, carryOut)
			dataForInsertion := currentResult % 10000
			carryOut = currentResult / 10000
			result.AddNodeAtTheBeginning(dataForInsertion)
			secondNumberPointer = secondNumberPointer.GetNextNode()
		}
		firstNumberPointer = firstNumberPointer.GetNextNode()
		currentOffset++
	}
	return &HPI{result}
}

func evaluateCurrentMultiplication(multiplyingPointer, multiplierPointer *linked_list.SimpleNode, carryOut int) int {
	return (multiplyingPointer.GetData().(int) * multiplierPointer.GetData().(int)) + carryOut
}
