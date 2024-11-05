package high_precision_integers

import "github.com/web3foru/computer-science-group-golang/internal/application/src/data_structures/linked_list"

func SetUpHighPrecisionInteger(n int) HighPrecisionInteger {
	auxiliaryRepresentation := n
	auxiliaryList := linked_list.NewLinkedList()
	for auxiliaryRepresentation >= 10000 {
		lessSignificantDigits := auxiliaryRepresentation % 10000
		auxiliaryList.AddNodeAtTheBeginning(lessSignificantDigits)
		auxiliaryRepresentation = auxiliaryRepresentation / 10000
	}
	auxiliaryList.AddNodeAtTheBeginning(auxiliaryRepresentation)
	return HighPrecisionInteger{NumberRepresentation: auxiliaryList}
}

func (hpi *HighPrecisionInteger) getLinkedList() *linked_list.SimpleLinkedList {
	return hpi.NumberRepresentation
}
