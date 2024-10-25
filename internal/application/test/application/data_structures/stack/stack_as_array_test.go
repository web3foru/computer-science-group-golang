package stack

import (
	"testing"
)

func TestStackAsArray(t *testing.T) {
	t.Run("create an empty stack", func(t *testing.T) {
		newStack := CreateAnEmptyStackAsArray()
		if !newStack.IsEmpty() {
			t.Errorf("new stack should be empty")
		}
	})

	t.Run("append a new element to stack", func(t *testing.T) {
		newStack := GetAStackWithOneElement()
		if newStack.IsEmpty() {
			t.Errorf("new stack should not be empty")
		}
	})

	t.Run("append a new element to stack", func(t *testing.T) {
		newStack := CreateAnEmptyStackAsArray()
		newElemnt := "un elemento Nuevo"
		newStack.Push(newElemnt)

		if newStack.IsEmpty() {
			t.Errorf("new stack should not be empty")
		}
	})

	t.Run("pop elements of a stack", func(t *testing.T) {
		newStack := GetAStackWithNElements(10)
		poppedElement := newStack.Pop()

		if poppedElement != 9 || newStack.Peek() != 8 {
			t.Errorf("Popped element must be the last")
		}
	})

	t.Run("pop elements of an empty stack", func(t *testing.T) {
		newStack := GetAStackWithNElements(0)
		poppedElement := newStack.Pop()

		if poppedElement != nil {
			t.Errorf("there must be a null elemente due to empty stack")
		}
	})
}
