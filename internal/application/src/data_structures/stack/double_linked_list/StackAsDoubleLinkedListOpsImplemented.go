package double_linked_list

import (
	"github.com/web3foru/computer-science-group-golang/internal/application/src/data_structures/double_linked_list"
)

func CreateNewStackAsDoubleLinkedList() StackAsDoubleLinkedList {
	elements := linked_list.NewDoublyLinkedList()
	return StackAsDoubleLinkedList{items: elements, size: 0}
}

func (stack *StackAsDoubleLinkedList) Push(element interface{}) {
	stack.items.AddNodeAtEnd(element)
	stack.size++
}

func (stack *StackAsDoubleLinkedList) Size() int {
	return stack.size
}

func (stack *StackAsDoubleLinkedList) Pop() interface{} {
	if stack.size == 0 {
		return nil
	} else {
		lastElement := stack.items.GetTail()
		stack.items.RemoveNodeAtEnd()
		stack.size--
		return lastElement.GetValue()
	}
}

func (stack *StackAsDoubleLinkedList) Peek() interface{} {
	if stack.size != 0 {
		return stack.items.GetTail().GetValue()
	} else {
		return nil
	}
}
