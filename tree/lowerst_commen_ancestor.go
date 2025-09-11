package tree

import "DSA_Go/linked_list/models"

var node *models.TreeNode

func LowestCommonAncestor(root *models.TreeNode, p *models.TreeNode, q *models.TreeNode) *models.TreeNode {

	dfs(root, p, q)
	return node

}

func dfs(root *models.TreeNode, p *models.TreeNode, q *models.TreeNode) bool {
	if root == nil {
		return false
	}

	node_is_p_or_q := root.Val == p.Val || root.Val == q.Val
	left_contain_p_or_q := dfs(root.Left, p, q)
	right_contain_p_or_q := dfs(root.Right, p, q)

	if (node_is_p_or_q && left_contain_p_or_q) || (node_is_p_or_q && right_contain_p_or_q) || (left_contain_p_or_q && right_contain_p_or_q) {
		node = root
	}

	return node_is_p_or_q || left_contain_p_or_q || right_contain_p_or_q
}
