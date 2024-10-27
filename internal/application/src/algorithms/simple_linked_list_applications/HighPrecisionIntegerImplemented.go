package simple_linked_list_applications

import "github.com/web3foru/computer-science-group-golang/internal/application/src/data_structures/linked_list"

func ReverseList(sll *linked_list.SimpleLinkedList) *linked_list.SimpleLinkedList {
	principleNodePointer := sll.GetFirstNode()
	lastNodePointer := sll.GetLastNode()
	for !sll.IsTheEndOfTheList(principleNodePointer) {
		innerPointer := lastNodePointer
		lastNodePointer = principleNodePointer
		principleNodePointer.UpdateNext(innerPointer)
	}
	return sll
}
