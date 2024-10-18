package queue

func CreateAnQueAsAnArray() QueueAsAnArray {
	var array []interface{}
	return QueueAsAnArray{
		items: array,
		size:  0,
	}
}

func (queue *QueueAsAnArray) Enqueue(value interface{}) bool {
	queue.size++
	queue.items = append(queue.items, value)
	return true
}

func (queue *QueueAsAnArray) Size() int {
	return queue.size
}

func (queue *QueueAsAnArray) IsEmpty() bool {
	if queue.size == 0 {
		return true
	} else {
		return false
	}
}

func (queue *QueueAsAnArray) Dequeue() interface{} {
	if queue.IsEmpty() {
		return nil
	} else {

		currentObject := queue.items[0]
		queue.size--
		queue.items = append(queue.items[:1], queue.items[2:]...)
		return currentObject
	}
}

func (queue *QueueAsAnArray) Peek() interface{} {
	return queue.items[queue.size-1]
}
