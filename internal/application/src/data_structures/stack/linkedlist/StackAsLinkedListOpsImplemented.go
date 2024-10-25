package linkedlist

import "github.com/web3foru/computer-science-group-golang/internal/application/src/data_structures/linked_list"

func CreateNewStackAsLinkedList() StackAsLinkedList {
	elements := linked_list.NewLinkedList()
	return StackAsLinkedList{items: elements, size: 0}
}

func (stack *StackAsLinkedList) Push(element interface{}) {
	stack.items.AddNodeAtTheEnd(element)
	stack.size++
}
func (stack *StackAsLinkedList) IsEmpty() bool {
	return stack.size == 0
}

func (stack *StackAsLinkedList) Pop() interface{} {
	if stack.IsEmpty() {
		return nil
	} else {
		lastElementOfStack := stack.items.ReturnLastNodeAndErase()
		stack.size--
		return lastElementOfStack.GetData()
	}
}
