package tree

import "DSA_Go/linked_list/models"

// tc : O(n)
// sc : O(n)
type TreeNodeValue struct {
	node  *models.TreeNode
	count int
}

func FindWidestBinaryLevel(root *models.TreeNode) int {
	if root == nil {
		return 0
	}

	queue := make([]TreeNodeValue, 0)
	treeNodeValue := TreeNodeValue{root, 0}
	queue = append(queue, treeNodeValue)
	max_width := 0

	for len(queue) > 0 {
		size := len(queue)

		left_most_index := queue[0].count
		right_most_index := queue[size-1].count

		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]

			if node.node.Left != nil {
				treeNode := TreeNodeValue{node.node.Left, 2*node.count + 1}
				queue = append(queue, treeNode)
			}

			if node.node.Right != nil {
				treeNode := TreeNodeValue{node.node.Right, 2*node.count + 2}
				queue = append(queue, treeNode)
			}
			// right_most_index = node.count
		}
		max_width = max(max_width, right_most_index-left_most_index+1)
	}

	return max_width
}
