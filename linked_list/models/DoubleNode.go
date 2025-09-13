package models

type DoubleNode struct {
	Val  int
	Next *DoubleNode
	Prev *DoubleNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

