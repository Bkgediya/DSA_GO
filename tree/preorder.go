package tree

import (
	"DSA_Go/linked_list/models"
)

var preorder_index int = 0
var inorder_indexes_map = make(map[int]int)

func BuildBinaryTree(preorder []int, inorder []int) *models.TreeNode {
	for i, num := range inorder {
		inorder_indexes_map[num] = i
	}

	return build_subtree(0, len(inorder)-1, preorder, inorder)
}

func build_subtree(left, right int, preorder, inorder []int) *models.TreeNode {

	if left > right {
		return nil
	}

	val := preorder[preorder_index]
	inorder_index := inorder_indexes_map[val]

	node := &models.TreeNode{
		Val: val,
	}
	preorder_index++

	node.Left = build_subtree(left, inorder_index-1, preorder, inorder)
	node.Right = build_subtree(inorder_index+1, right, preorder, inorder)

	return node
}
