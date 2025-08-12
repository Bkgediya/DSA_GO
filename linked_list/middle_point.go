package linkedlist

import "DSA_Go/linked_list/models"

// FindListMiddlePoint find the middle point of a linked list
// Time Complexity: O(n)
// Space Complexity: O(1)
func FindListMiddlePoint(head *models.ListNode) *models.ListNode {
	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}
