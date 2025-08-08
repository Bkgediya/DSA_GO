package linkedlist

import "DSA_Go/linked_list/models"

// time: O(n), space: O(1)
func ReverseLinkedList(head *models.ListNode) *models.ListNode {
	if head == nil {
		return nil
	}

	curr_node, prev_node := head, (*models.ListNode)(nil)

	for curr_node != nil {
		next_node := curr_node.Next
		curr_node.Next = prev_node
		prev_node = curr_node
		curr_node = next_node
	}

	return prev_node
}
