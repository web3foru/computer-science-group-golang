package queue

type QueueOperations interface {
	Enqueue(data interface{}) bool
	Dequeue() interface{}
	Peek() interface{}
	Size() int
	IsEmpty() bool
}
