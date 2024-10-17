package stack

import (
	"github.com/web3foru/computer-science-group-golang/internal/application/src/data_structures/stack/array"
)

func CreateAnEmptyStackAsArray() array.StackAsAnArray {
	return array.CreateNewStackAsArray()
}

func GetAStackWithOneElement() array.StackAsAnArray {
	return GetAStackWithNElements(1)
}

func GetAStackWithNElements(quantity int) array.StackAsAnArray {
	newStack := array.CreateNewStackAsArray()
	for i := 0; i < quantity; i++ {
		newStack.Push(i)
	}
	return newStack
}
