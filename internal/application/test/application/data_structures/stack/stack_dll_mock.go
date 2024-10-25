package stack

import "github.com/web3foru/computer-science-group-golang/internal/application/src/data_structures/stack/double_linked_list"

func CreateAnEmptyStack() double_linked_list.StackAsDoubleLinkedList {
	return double_linked_list.CreateNewStackAsDoubleLinkedList()
}

func CreateAStackWithOneItem() double_linked_list.StackAsDoubleLinkedList {
	return CreateNewStackWithNItems(1)
}

func CreateNewStackWithNItems(n int) double_linked_list.StackAsDoubleLinkedList {
	newStack := CreateAnEmptyStack()
	for i := 0; i < n; i++ {
		newStack.Push(i)
	}
	return newStack
}
