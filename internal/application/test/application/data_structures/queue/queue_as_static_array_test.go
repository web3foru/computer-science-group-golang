package queue

import (
	"github.com/web3foru/computer-science-group-golang/internal/application/src/data_structures/queue/implementations/array"
	"testing"
)

func TestQueueAsArray(t *testing.T) {
	t.Run("Create an empty queue", func(t *testing.T) {
		newQueue := array.CreateAQueAsAnArray()
		if !newQueue.IsEmpty() {
			t.Errorf("New queue should be empty")
		}
	})
}
