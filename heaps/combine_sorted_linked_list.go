package heaps

import (
	"DSA_Go/linked_list/models"
	"container/heap"
)

func ComibineSortedLinkedList(lists []*models.ListNode) *models.ListNode {
	minListNodeHeap := NewMinHeapListNode()

	// Add all the head nodes to the heap
	for _, node := range lists {
		if node != nil {
			heap.Push(minListNodeHeap, node)
		}
	}

	dummy := &models.ListNode{}
	curr := dummy

	for minListNodeHeap.Len() > 0 {
		// Pop smallest node
		node := heap.Pop(minListNodeHeap).(*models.ListNode)
		curr.Next = node
		curr = curr.Next

		// If this list still has nodes, push next into heap
		if node.Next != nil {
			heap.Push(minListNodeHeap, node.Next)
		}
	}

	return dummy.Next
}
