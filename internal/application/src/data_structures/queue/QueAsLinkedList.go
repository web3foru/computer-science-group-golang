package queue

import "github.com/web3foru/computer-science-group-golang/internal/application/src/data_structures/linked_list"

func CreateQueueAsSimpleLinkedList() QueueStructure {
	linkedListItems := linked_list.NewLinkedList()
	return QueueStructure{items: linkedListItems, size: 0}
}
