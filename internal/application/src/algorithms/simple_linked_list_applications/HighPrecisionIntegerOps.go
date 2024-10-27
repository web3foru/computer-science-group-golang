package simple_linked_list_applications

import "github.com/web3foru/computer-science-group-golang/internal/application/src/data_structures/linked_list"

type HighPrecisionIntegerOps interface {
	reverseList() linked_list.SimpleLinkedList
	sum(number1 HighPrecisionInteger, number2 HighPrecisionInteger) HighPrecisionIntegerOps
}
