package heaps

import (
	"DSA_Go/linked_list/models"
	"container/heap"
)

type MinListNodeHeap []*models.ListNode

func (h MinListNodeHeap) Len() int           { return len(h) }
func (h MinListNodeHeap) Less(i, j int) bool { return h[i].Val < h[j].Val }
func (h MinListNodeHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinListNodeHeap) Push(x interface{}) {
	*h = append(*h, x.(*models.ListNode))
}

func (h *MinListNodeHeap) Pop() interface{} {
	old := *h
	n := len(old)
	node := old[n-1]
	*h = old[:n-1]
	return node
}

func NewMinHeapListNode() *MinListNodeHeap {
	h := &MinListNodeHeap{}
	heap.Init(h)
	return h
}
