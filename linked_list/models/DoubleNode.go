package models

type DoubleNode struct {
	Val  int
	Next *DoubleNode
	Prev *DoubleNode
}

type GraphNode struct {
	Val       int
	Neighbors []*GraphNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
