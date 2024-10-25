package stack

import "github.com/web3foru/computer-science-group-golang/internal/application/src/data_structures/stack/linkedlist"

func CreateAStackWithNElements(size int) linkedlist.StackAsLinkedList {
	newStack := linkedlist.CreateNewStackAsLinkedList()
	for i := 0; i < size; i++ {
		newStack.Push(i)
	}
	return newStack
}
