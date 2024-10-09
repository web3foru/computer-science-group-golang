package queue

import (
	"fmt"
	"github.com/web3foru/computer-science-group-golang/internal/application/src/data_structures/linked_list"
)

func (queue *QueueStructure) Enqueue(value interface{}) bool {
	newNode := linked_list.CreateNewNode(value)
	queue.items.AddNodeAtTheEnd(newNode)
	queue.size++
	return true
}

func (queue *QueueStructure) Size() int {
	return queue.size
}

func (queue *QueueStructure) Dequeue() interface{} {
	if queue.IsEmpty() {
		return nil
	} else {
		queue.size--
		return queue.items.ReturnFirstNodeAndErase().String()
	}
}

func (queue *QueueStructure) Peek() interface{} {
	if queue.IsEmpty() {
		return nil
	} else {
		return queue.items.GetFirstNode().String()
	}
}

func (queue *QueueStructure) IsEmpty() bool {
	if queue.Size() == 0 {
		fmt.Sprintln("The Queue is empty")
		return true
	}
	return false
}
