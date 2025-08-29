package heaps

import "container/heap"

func FindKMostFrequentString(words []string, k int) []string {
	freq_of_words := make(map[string]int)

	for _, word := range words {
		freq_of_words[word]++
	}

	// build MaxHeap
	h := NewMaxHeap()
	for word, freq := range freq_of_words {
		heap.Push(h, Item{Word: word, Count: freq})
	}

	// extract top k
	result := []string{}
	for i := 0; i < k && h.Len() > 0; i++ {
		item := heap.Pop(h).(Item)
		result = append(result, item.Word)
	}
	return result
}

func FindKMostFrequentStringWithMinHeap(words []string, k int) []string {
	freq_of_words := make(map[string]int)

	for _, word := range words {
		freq_of_words[word]++
	}

	h := NewMinHeap()
	for word, freq := range freq_of_words {
		heap.Push(h, MinItem{Word: word, Count: freq})
	}
	result := []string{}
	n := h.Len()

	for i := n - 1; i < (n-k) && h.Len() > 0; i++ {
		item := heap.Pop(h).(MinItem)
		result = append(result, item.Word)
	}

	return result
}
