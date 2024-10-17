package queue

import "github.com/web3foru/computer-science-group-golang/internal/application/src/data_structures/queue"

func GetAnEmptyQueue() queue.QueueAsLinkedList {
	return queue.CreateQueueAsSimpleLinkedList()
}

func GetAQueueWithTwoItems() queue.QueueAsLinkedList {
	return GetAQueueWithNItems(2)
}

func GetAQueueWithNItems(numberOfItems int) queue.QueueAsLinkedList {
	newQueue := GetAnEmptyQueue()
	for i := 0; i < numberOfItems; i++ {
		newQueue.Enqueue(i)
	}
	return newQueue
}
