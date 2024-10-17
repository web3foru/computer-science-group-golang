package queue

import (
	"github.com/web3foru/computer-science-group-golang/internal/application/src/data_structures/queue/linkedlist"
)

func GetAnEmptyQueue() linkedlist.QueueAsLinkedList {
	return linkedlist.CreateQueueAsSimpleLinkedList()
}

func GetAQueueWithTwoItems() linkedlist.QueueAsLinkedList {
	return GetAQueueWithNItems(2)
}

func GetAQueueWithNItems(numberOfItems int) linkedlist.QueueAsLinkedList {
	newQueue := GetAnEmptyQueue()
	for i := 0; i < numberOfItems; i++ {
		newQueue.Enqueue(i)
	}
	return newQueue
}
