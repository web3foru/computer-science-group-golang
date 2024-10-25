package stack

import (
	"testing"
)

func TestStackAsDoubleLinkedList(t *testing.T) {
	t.Run("Create an empty linked list", func(t *testing.T) {
		stackAsDoubledLinkedList := CreateAnEmptyStack()
		if stackAsDoubledLinkedList.Size() != 0 {
			t.Error("must create a new empty stack")
		}
	})

	t.Run("Create an linked list with one item", func(t *testing.T) {
		stackWithOneItem := CreateAStackWithOneItem()
		if stackWithOneItem.Size() != 1 {
			t.Error("must create a new stack with one item")
		}
	})

	t.Run("Pop an element from an stack", func(t *testing.T) {
		stack := CreateNewStackWithNItems(10)
		poppedElement := stack.Pop()
		if poppedElement != 9 {
			t.Error("must pop an element from a stack")
		}
	})

	t.Run("Peek an element from an stack", func(t *testing.T) {
		stack := CreateNewStackWithNItems(10)
		peekedElement := stack.Peek()
		if peekedElement != 9 {
			t.Error("must peek an element from a stack")
		}
	})

	t.Run("pop an element from an empty stack, must be null", func(t *testing.T) {
		stack := CreateAnEmptyStack()
		poppedElement := stack.Pop()
		if poppedElement != nil {
			t.Error("must pop an empty stack")
		}
	})

	t.Run("peek an element from an empty stack, must be null", func(t *testing.T) {
		emptyStack := CreateAnEmptyStack()
		peekedElement := emptyStack.Peek()
		if peekedElement != nil {
			t.Error("must peek an empty stack")
		}
	})
}
