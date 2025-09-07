package tree

import "DSA_Go/linked_list/models"

// tc : O(n)
// sc : O(n)
func FindRightMostNodeInBinaryTree(root *models.TreeNode) []int {
	result := []int{}
	queue := []*models.TreeNode{root}

	for len(queue) > 0 {
		size := len(queue)

		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]

			if i == size-1 {
				result = append(result, node.Val)
			}

			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}

	return result
}
