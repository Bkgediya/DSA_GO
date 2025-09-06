package tree

import "DSA_Go/linked_list/models"

func InvertBinaryTree(root *models.TreeNode) *models.TreeNode {
	if root == nil {
		return nil
	}
	stack := []*models.TreeNode{root}

	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		node.Left, node.Right = node.Right, node.Left

		if node.Left != nil {
			stack = append(stack, node.Left)
		}

		if node.Right != nil {
			stack = append(stack, node.Right)
		}
	}

	return root
}
