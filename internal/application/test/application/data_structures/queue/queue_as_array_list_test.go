package queue

import "testing"

func TestNewQueue(t *testing.T) {
	t.Run("Create a new queue based on linkedList", func(t *testing.T) {
		newQueue := GetAnEmptyQueue()
		if newQueue.Size() != 0 {
			t.Error("Queue size is not zero")
		}
	})

	t.Run("Enqueue a new item to the queue", func(t *testing.T) {
		newQueue := GetAQueueWithTwoItems()
		if newQueue.Size() != 2 {
			t.Error("Queue size is not two")
		}
	})

	t.Run("Dequeue an item from the queue", func(t *testing.T) {
		newQueue := GetAQueueWithNItems(10)
		dequeuedItem := newQueue.Dequeue()
		if dequeuedItem != "0" {
			t.Error("Dequeued item is not zero")
		}
	})

	t.Run("Dequeue an empty queue", func(t *testing.T) {
		newQueue := GetAnEmptyQueue()
		dequeuedItem := newQueue.Dequeue()
		if dequeuedItem != nil {
			t.Error("queue must be empty")
		}
	})

	t.Run("peek an element of the queue", func(t *testing.T) {
		newQueue := GetAQueueWithNItems(10)
		dequeuedItem := newQueue.Peek()
		if dequeuedItem == nil {
			t.Error("queue must has an item")
		}
	})
}
