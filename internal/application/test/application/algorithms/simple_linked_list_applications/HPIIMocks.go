package simple_linked_list_applications

import (
	"github.com/web3foru/computer-science-group-golang/internal/application/src/data_structures/linked_list"
	"github.com/web3foru/computer-science-group-golang/internal/application/test/application/data_structures/simple_linked_list"
)

func GetAListForReversing() *linked_list.SimpleLinkedList {
	return simple_linked_list.GetAnSimpleListWithNNodes(5)
}

func GetAListWithTwoNodesForSwaping() *linked_list.SimpleLinkedList {
	return simple_linked_list.GetAnSimpleListWithNNodes(2)
}
