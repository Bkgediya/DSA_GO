package models

type LRUCache struct {
	capacity int
	head     *ListNode
	tail     *ListNode
	cache    map[int]*ListNode
}
