package queue

import "github.com/web3foru/computer-science-group-golang/internal/application/src/data_structures/queue"

func GetAnEmptyQueue() queue.QueueStructure {
	return queue.CreateQueueAsSimpleLinkedList()
}

func GetAQueueWithTwoItems() queue.QueueStructure {
	return GetAQueueWithNItems(2)
}

func GetAQueueWithNItems(numberOfItems int) queue.QueueStructure {
	newQueue := GetAnEmptyQueue()
	for i := 0; i < numberOfItems; i++ {
		newQueue.Enqueue(i)
	}
	return newQueue
}
