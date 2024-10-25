package queue

import (
	linkedlist2 "github.com/web3foru/computer-science-group-golang/internal/application/src/data_structures/queue/implementations/linkedlist"
)

func GetAnEmptyQueue() linkedlist2.QueueAsLinkedList {
	return linkedlist2.CreateQueueAsSimpleLinkedList()
}

func GetAQueueWithTwoItems() linkedlist2.QueueAsLinkedList {
	return GetAQueueWithNItems(2)
}

func GetAQueueWithNItems(numberOfItems int) linkedlist2.QueueAsLinkedList {
	newQueue := GetAnEmptyQueue()
	for i := 0; i < numberOfItems; i++ {
		newQueue.Enqueue(i)
	}
	return newQueue
}
