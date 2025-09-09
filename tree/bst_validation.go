package tree

import (
	"DSA_Go/linked_list/models"
	"math"
)

func BinarySearchTreeValidation(root *models.TreeNode) bool {
	return is_within_bound(root, math.Inf(-1), math.Inf(1))
}

func is_within_bound(node *models.TreeNode, lower_bound, upper_bound float64) bool {
	if node == nil {
		return true
	}

	if !(lower_bound < float64(node.Val) && float64(node.Val) < upper_bound) {
		return false
	}

	if !is_within_bound(node.Left, lower_bound, float64(node.Val)) {
		return false
	}

	return is_within_bound(node.Right, float64(node.Val), upper_bound)
}
