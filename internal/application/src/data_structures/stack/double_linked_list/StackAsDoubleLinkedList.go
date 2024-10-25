package double_linked_list

import (
	"github.com/web3foru/computer-science-group-golang/internal/application/src/data_structures/double_linked_list"
)

type StackAsDoubleLinkedList struct {
	items *linked_list.DoublyLinkedList
	size  int
}
