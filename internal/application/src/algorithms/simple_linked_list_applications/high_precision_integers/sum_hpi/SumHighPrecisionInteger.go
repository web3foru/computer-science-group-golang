package sum_hpi

import (
	"github.com/web3foru/computer-science-group-golang/internal/application/src/algorithms/simple_linked_list_applications/high_precision_integers"
	"github.com/web3foru/computer-science-group-golang/internal/application/src/data_structures/linked_list"
)

type SumHighPrecisionInteger interface {
	Sum(number1 *high_precision_integers.HighPrecisionInteger, number2 *high_precision_integers.HighPrecisionInteger) high_precision_integers.HighPrecisionInteger
	addCarryOut(result *linked_list.SimpleLinkedList, carryOut *linked_list.SimpleNode)
	evaluateRemainPartOfTheNumber(number *high_precision_integers.HighPrecisionInteger, checkpointNode *linked_list.SimpleNode, carryOut *linked_list.SimpleNode, currentResult *linked_list.SimpleLinkedList) *linked_list.SimpleNode
	evaluateSum(node1 *linked_list.SimpleNode, node2 *linked_list.SimpleNode, carryOut *linked_list.SimpleNode) (*linked_list.SimpleNode, int)
}
