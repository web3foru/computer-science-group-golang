package simple_linked_list_applications

import "github.com/web3foru/computer-science-group-golang/internal/application/src/data_structures/linked_list"

type HighPrecisionInteger struct {
	numberRepresentation *linked_list.SimpleLinkedList
}

func (number1 *HighPrecisionInteger) GetLinkedList() *linked_list.SimpleLinkedList {
	return number1.numberRepresentation
}
