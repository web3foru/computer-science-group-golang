package linkedlist

import "github.com/web3foru/computer-science-group-golang/internal/application/src/data_structures/linked_list"

type StackAsLinkedList struct {
	items *linked_list.SimpleLinkedList
	size  int
}
