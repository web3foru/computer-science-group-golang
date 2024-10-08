package simple_linked_list

import "github.com/web3foru/computer-science-group-golang/internal/application/src/data_structures/linked_list"

func GetAnEmptySimpleList() *linked_list.SimpleLinkedList {
	return linked_list.NewLinkedList()
}

func GetAnSimpleListWithOneNode() *linked_list.SimpleLinkedList {
	newLinked := linked_list.NewLinkedList()
	newLinked.AddNodeAtTheBeginning("a new node")
	return newLinked
}

func GetAnSimpleListWithNNodes(nodeQuantity int) *linked_list.SimpleLinkedList {
	linkedList := GetAnEmptySimpleList()
	for i := 0; i < nodeQuantity; i++ {
		linkedList.AddNodeAtTheEnd(i)
	}
	return linkedList
}
