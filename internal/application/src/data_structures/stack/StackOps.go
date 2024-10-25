package stack

type StackOps interface {
	Push(interface{})
	IsEmpty() bool
	Pop() interface{}
	Peek() interface{}
}
