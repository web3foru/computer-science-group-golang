package simple_linked_list_applications

import "github.com/web3foru/computer-science-group-golang/internal/application/src/data_structures/linked_list"

func SetUpHighPrecisionInteger(n int) HighPrecisionInteger {
	auxiliaryRepresentation := n
	auxiliaryList := linked_list.NewLinkedList()

	if isShortNumber(n) {
		auxiliaryList = setUpAShortHPI(n)
	} else {
		for auxiliaryRepresentation >= 10000 {
			lessSignificantDigits := auxiliaryRepresentation % 10000
			auxiliaryList.AddNodeAtTheBeginning(lessSignificantDigits)
			auxiliaryRepresentation = auxiliaryRepresentation / 10000
		}
		auxiliaryList.AddNodeAtTheBeginning(auxiliaryRepresentation)
	}
	return HighPrecisionInteger{numberRepresentation: auxiliaryList}
}

func setUpAShortHPI(n int) *linked_list.SimpleLinkedList {
	newsll := linked_list.NewLinkedList()
	newsll.AddNodeAtTheBeginning(n)
	return newsll
}

func isShortNumber(n int) bool {
	if n < 10000 {
		return true
	} else {
		return false
	}
}

func Sum(hpiA HighPrecisionInteger, hpib HighPrecisionInteger) HighPrecisionInteger {
	return SetUpHighPrecisionInteger(1)
}

func Difference(hpiA HighPrecisionInteger, hpib HighPrecisionInteger) HighPrecisionInteger {
	return SetUpHighPrecisionInteger(1)
}

func (hpi *HighPrecisionInteger) GetLinkedList() *linked_list.SimpleLinkedList {
	return hpi.numberRepresentation
}
