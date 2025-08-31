package heaps

import "container/heap"

// time: O(log n)
// space: O(n)
type MinMaxHeap struct {
	low  MaxIntHeap // low half
	high MinIntHeap // upper half
}

func NewMedianFinder() *MinMaxHeap {
	h := &MinMaxHeap{}
	heap.Init(&h.low)
	heap.Init(&h.high)
	return h
}

func (mmh *MinMaxHeap) Add(num int) {
	if mmh.low.Len() == 0 || num < mmh.low[0] {
		heap.Push(&mmh.low, num)
	} else {
		heap.Push(&mmh.high, num)
	}

	// rebalancing
	if mmh.low.Len() > mmh.high.Len()+1 {
		heap.Push(&mmh.high, heap.Pop(&mmh.low))
	} else if mmh.high.Len() > mmh.low.Len() {
		heap.Push(&mmh.low, heap.Pop(&mmh.high))
	}
}

func (mmh *MinMaxHeap) FindMedian() float64 {
	total := mmh.low.Len() + mmh.high.Len()

	if total == 0 {
		return 0
	}

	if total%2 == 0 {
		return float64(mmh.low[0]+mmh.high[0]) / 2.0
	}
	return float64(mmh.low[0])
}
