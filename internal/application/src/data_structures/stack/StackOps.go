package stack

type StackOps interface {
	Pop() interface{}
	Push(interface{})
	Peek() interface{}
}
