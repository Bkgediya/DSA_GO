package tree

import "DSA_Go/linked_list/models"

//tc : O(n)
//sc : O(n)
func BalancedBinaryTreeValidation(root *models.TreeNode) bool {
	return get_height_of_tree(root) != -1
}

func get_height_of_tree(root *models.TreeNode) int {
	if root == nil {
		return 0
	}

	left_height := get_height_of_tree(root.Left)
	right_height := get_height_of_tree(root.Right)

	if left_height == -1 || right_height == -1 {
		return -1
	}

	if abs(left_height-right_height) > 1 {
		return -1
	}

	return (1 + max(left_height, right_height))
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
