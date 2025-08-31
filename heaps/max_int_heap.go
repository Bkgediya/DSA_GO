package heaps

type MaxIntHeap []int

// len
func (h MaxIntHeap) Len() int {
	return len(h)
}

// less
func (h MaxIntHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

// swap
func (h MaxIntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MaxIntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MaxIntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]

	return x
}
