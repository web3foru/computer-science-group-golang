package array

func CreateNewStackAsArray() StackAsAnArray {
	var array []interface{}
	return StackAsAnArray{items: array, size: 0}
}

func (stack *StackAsAnArray) Push(value interface{}) {
	stack.size += 1
	stack.items = append(stack.items, value)
}

func (stack *StackAsAnArray) IsEmpty() bool {
	return stack.size == 0
}

func (stack *StackAsAnArray) Pop() interface{} {
	if stack.IsEmpty() {
		return nil
	} else {
		currentItem := stack.items[stack.size-1]
		stack.items = stack.items[:stack.size-1]
		stack.size -= 1
		return currentItem
	}
}
func (stack *StackAsAnArray) Peek() interface{} {
	return stack.items[len(stack.items)-1]
}
