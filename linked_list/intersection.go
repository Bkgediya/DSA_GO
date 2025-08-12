package linkedlist

import (
	"DSA_Go/linked_list/models"
)

//find the intersection of the linkedlist
// time complexity: O(n)
// space complexity: O(1)

func FindTheIntersection(head1 *models.ListNode, head2 *models.ListNode) *models.ListNode {
	if head1 == nil || head2 == nil {
		return nil
	}
	ptrA, ptrB := head1, head2

	for ptrA != ptrB {
		if ptrA == nil {
			ptrA = head2
		} else {
			ptrA = ptrA.Next
		}

		if ptrB == nil {
			ptrB = head1
		} else {
			ptrB = ptrB.Next
		}
	}

	return ptrA
}
