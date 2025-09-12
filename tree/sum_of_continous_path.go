package tree

import (
	"DSA_Go/linked_list/models"
	"math"
)

var max_sum = int(math.Inf(-1))

func FindContinousPathInTree(root *models.TreeNode) int {
	max_find_helper(root)
	return max_sum
}

func max_find_helper(root *models.TreeNode) int {
	if root == nil {
		return 0
	}
	left_sum := max(max_find_helper(root.Left), 0)
	right_sum := max(max_find_helper(root.Right), 0)

	max_sum = max(max_sum, root.Val+left_sum+right_sum)

	return root.Val + max(left_sum, right_sum)
}
