package high_precision_integers

import "github.com/web3foru/computer-science-group-golang/internal/application/src/data_structures/linked_list"

type HighPrecisionInteger struct {
	NumberRepresentation *linked_list.SimpleLinkedList
}

func (number1 *HighPrecisionInteger) GetLinkedList() *linked_list.SimpleLinkedList {
	return number1.NumberRepresentation
}
