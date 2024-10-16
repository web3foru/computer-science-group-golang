package queue

import (
	"fmt"
	"github.com/web3foru/computer-science-group-golang/internal/application/src/data_structures/linked_list"
)

func CreateQueueAsSimpleLinkedList() QueueAsLinkedList {
	linkedListItems := linked_list.NewLinkedList()
	return QueueAsLinkedList{items: linkedListItems, size: 0}
}

func (queue *QueueAsLinkedList) Enqueue(value interface{}) bool {
	queue.items.AddNodeAtTheEnd(value)
	queue.size++
	return true
}

func (queue *QueueAsLinkedList) Size() int {
	return queue.size
}

func (queue *QueueAsLinkedList) Dequeue() interface{} {
	if queue.IsEmpty() {
		return nil
	} else {
		queue.size--
		return queue.items.ReturnFirstNodeAndErase().String()
	}
}

func (queue *QueueAsLinkedList) Peek() interface{} {
	if queue.IsEmpty() {
		return nil
	} else {
		return queue.items.GetFirstNode().String()
	}
}

func (queue *QueueAsLinkedList) IsEmpty() bool {
	if queue.Size() == 0 {
		fmt.Sprintln("The Queue is empty")
		return true
	}
	return false
}
