package linkedlist

import (
	"DSA_Go/linked_list/models"
	"fmt"
)

func RemoveNthNodeFromEnd(head *models.ListNode, n int) (*models.ListNode, error) {
	if head == nil || n < 1 {
		return nil, fmt.Errorf("empty list or invalid n")
	}

	dummy := &models.ListNode{Next: head}
	tailor := dummy
	leader := dummy

	for i := 0; i < n; i++ {
		leader = leader.Next
		if leader == nil {
			return nil, fmt.Errorf("invalid n")
		}
	}

	for leader != nil && leader.Next != nil {
		tailor = tailor.Next
		leader = leader.Next
	}

	tailor.Next = tailor.Next.Next
	return dummy.Next, nil
}
