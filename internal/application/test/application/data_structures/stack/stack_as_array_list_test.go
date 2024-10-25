package stack

import "testing"

func TestStackArrayList(t *testing.T) {
	t.Run("Create an empty stack ", func(t *testing.T) {
		newStack := CreateAStackWithNElements(0)
		if !newStack.IsEmpty() {
			t.Errorf("Stack must be empty initialized")
		}
	})

	t.Run("Append Element to an empty stack", func(t *testing.T) {
		newStack := CreateAStackWithNElements(1)
		if newStack.IsEmpty() {
			t.Errorf("Stack must be not empty")
		}
	})

	t.Run("Pop an element from stack", func(t *testing.T) {
		newStack := CreateAStackWithNElements(2)
		poppedElement := newStack.Pop()
		if poppedElement != 1 {
			t.Errorf("popped element must be the last one")
		}
	})

	t.Run("Pop an element from large stack", func(t *testing.T) {
		newStack := CreateAStackWithNElements(100)
		poppedElement := newStack.Pop()
		if poppedElement != 99 {
			t.Errorf("popped element must be the last one")
		}
	})

}
