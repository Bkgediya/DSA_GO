package tree

import (
	"DSA_Go/linked_list/models"
	"fmt"
)

func InOrderTraversalOfTree(root *models.TreeNode) {
	if root == nil {
		return
	}

	InOrderTraversalOfTree(root.Left)
	fmt.Printf("%d ", root.Val)
	InOrderTraversalOfTree(root.Right)
}
