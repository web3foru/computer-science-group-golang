package difference_hpi

import (
	"github.com/web3foru/computer-science-group-golang/internal/application/src/algorithms/simple_linked_list_applications/high_precision_integers"
	"github.com/web3foru/computer-science-group-golang/internal/application/src/data_structures/linked_list"
)

type SubtractionStructure struct {
	minuend           *high_precision_integers.HighPrecisionInteger
	subtrahend        *high_precision_integers.HighPrecisionInteger
	minuendPointer    *linked_list.SimpleNode
	subtrahendPointer *linked_list.SimpleNode
	borrow            int
	result            *linked_list.SimpleLinkedList
}
