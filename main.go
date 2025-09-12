package main

import (
	"DSA_Go/arrays"
	binarysearch "DSA_Go/binary_search"
	hashmapsets "DSA_Go/hashmap_sets"
	"DSA_Go/heaps"
	"DSA_Go/intervals"
	linkedlist "DSA_Go/linked_list"
	"DSA_Go/linked_list/models"
	prefixsums "DSA_Go/prefix_sums"
	"DSA_Go/sliding_window"
	"DSA_Go/stack"
	find_most_frequent "DSA_Go/test"
	"DSA_Go/tree"
	"fmt"
)

func main() {
	nums := []int{-5, -2, 3, 4, 6}
	target := 7

	value := arrays.TwoSum(nums, target)
	fmt.Printf("%+v", value)

	nums = []int{0, -1, 2, -3, 1}
	triplet := arrays.TripletSum(nums)
	fmt.Printf("%+v", triplet)

	val_palindrome := "21.0.2.2021"

	is_palindrome := arrays.Palindrome(val_palindrome)
	fmt.Printf("%+v", is_palindrome)

	nums_container := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	max_water := arrays.LargestContainer(nums_container)
	fmt.Printf("%+v", max_water)

	nums_pair := []int{-1, 3, 4, 2}
	target = 3
	pairs := hashmapsets.PairSum(nums_pair, target)
	fmt.Printf("%+v", pairs)
	fmt.Println()
	grid_arr := [][]int{
		{5, 3, 2, 6, 7, 8, 9, 1, 2},
		{6, 7, 2, 1, 9, 5, 3, 4, 8},
		{1, 9, 8, 3, 4, 2, 5, 6, 7},
		{8, 5, 9, 7, 6, 1, 4, 2, 3},
		{4, 2, 6, 8, 5, 3, 7, 9, 1},
		{7, 1, 3, 9, 2, 4, 8, 5, 6},
		{9, 6, 1, 5, 3, 7, 2, 8, 4},
		{2, 8, 7, 4, 1, 9, 6, 3, 5},
		{3, 4, 5, 2, 8, 6, 1, 7, 9},
	}
	is_valid_sudoku := hashmapsets.ValidateSudoku(grid_arr)
	fmt.Printf("%+v", is_valid_sudoku)

	fmt.Println()

	striping_arr := [][]int{
		{1, 2, 3, 4, 5},
		{6, 0, 8, 9, 10},
		{11, 12, 13, 14, 15},
		{16, 17, 18, 19, 0},
	}
	zero_stripped_arr := hashmapsets.ZeroStriping(striping_arr)
	fmt.Printf("%+v\n", zero_stripped_arr)

	head := &models.ListNode{Val: 1}
	head.Next = &models.ListNode{Val: 2}
	head.Next.Next = &models.ListNode{Val: 3}

	reverse_list := linkedlist.ReverseLinkedList(head)

	for node := reverse_list; node != nil; node = node.Next {
		fmt.Printf("%+v", node.Val)
	}

	head = &models.ListNode{Val: 1}
	head.Next = &models.ListNode{Val: 2}
	head.Next.Next = &models.ListNode{Val: 4}
	head.Next.Next.Next = &models.ListNode{Val: 7}
	head.Next.Next.Next.Next = &models.ListNode{Val: 3}
	head.Next.Next.Next.Next.Next = &models.ListNode{Val: 10}

	head, err := linkedlist.RemoveNthNodeFromEnd(head, 3)
	if err != nil {
		fmt.Printf("%+v", err)
	}

	for node := head; node != nil; node = node.Next {
		fmt.Printf("%+v\n", node.Val)
	}

	fmt.Println()
	fmt.Println()

	lru := linkedlist.Constructor(2)

	lru.Put(1, 1)
	lru.Put(2, 2)
	fmt.Println(lru.Get(1))

	lru.Put(3, 3)
	fmt.Println(lru.Get(2))

	lru.Put(4, 4)
	fmt.Println(lru.Get(1))
	fmt.Println(lru.Get(3))
	fmt.Println(lru.Get(4))

	// detect cuycle in linkedlist
	head = &models.ListNode{Val: 1}
	head.Next = &models.ListNode{Val: 2}
	head.Next.Next = &models.ListNode{Val: 3}
	head.Next.Next.Next = &models.ListNode{Val: 4}
	head.Next.Next.Next.Next = head.Next

	has_cycle := linkedlist.DetectCycle(head)
	fmt.Printf("%+v\n", has_cycle)

	head = &models.ListNode{Val: 1}
	head.Next = &models.ListNode{Val: 2}
	head.Next.Next = &models.ListNode{Val: 3}
	head.Next.Next.Next = &models.ListNode{Val: 4}
	head.Next.Next.Next.Next = &models.ListNode{Val: 5}
	head.Next.Next.Next.Next.Next = &models.ListNode{Val: 6}

	mid_node := linkedlist.FindListMiddlePoint(head)
	fmt.Println(mid_node.Val)

	tailNode := &models.ListNode{Val: 8}
	tailNode.Next = &models.ListNode{Val: 10}

	head = &models.ListNode{Val: 1}
	head.Next = &models.ListNode{Val: 2}
	head.Next.Next = &models.ListNode{Val: 3}
	head.Next.Next.Next = tailNode
	head2 := &models.ListNode{Val: 4}
	head2.Next = &models.ListNode{Val: 5}
	head2.Next.Next = tailNode

	// find the happy number
	happy_num := linkedlist.IsHappyNumber(116)
	fmt.Printf("%+v\n", happy_num)

	is_anagram := sliding_window.FindSlidingWindow("caabab", "aba")
	fmt.Printf("sliding window count is %+v\n", is_anagram)

	subString_length := sliding_window.LongestSubstring("cabcdeca")
	fmt.Printf("sliding window count is %+v\n", subString_length)

	// longest uniform substring
	longest_uniform_substring := sliding_window.FindLongestUniformSubstring("aabcdcca", 2)
	fmt.Printf("longest uniform substring is %+v\n", longest_uniform_substring)

	insertion_array := []int{1, 2, 4, 5, 7, 8, 9}
	insertion_index := binarysearch.FindInsertionIndex(insertion_array, 6)
	fmt.Printf("insertion index is %+v\n", insertion_index)

	occurance_array := []int{1, 2, 3, 4, 4, 4, 5, 6, 7, 8, 9, 10, 11}
	occurance := binarysearch.FirstLastOccurance(occurance_array, 4)
	fmt.Printf("occurance is %+v\n", occurance)

	// cut enough wood
	wood_array := []int{2, 6, 3, 8}
	wood_cut := binarysearch.CuttingWood(wood_array, 7)
	fmt.Printf("wood cut is %+v\n", wood_cut)

	// find target in rotated array
	rotated_array := []int{8, 9, 1, 2, 3, 4, 5, 6, 7}
	target_index := binarysearch.FindTheTargetInRotatedArray(rotated_array, 5)
	fmt.Printf("target index is %+v\n", target_index)

	// valid parenthesis

	valid_parenthesis := stack.IsValidParenthesis("()[][{[}]]")
	fmt.Printf("valid parenthesis is %+v\n", valid_parenthesis)

	// find the next large number to right
	next_large_number := stack.FindNextLargeNumberToRight([]int{5, 2, 4, 6, 1})
	fmt.Printf("next large number is %+v\n", next_large_number)

	// evaluate expression
	evaluate_expression := stack.EvaluateExpression("18-(7+(2-4))")
	fmt.Printf("evaluate expression is %+v\n", evaluate_expression)

	// top k frequent string max heap
	arr_str := []string{"go", "coding", "byte", "byte", "go", "interview", "go"}
	top_k_frequent := heaps.FindKMostFrequentString(arr_str, 2)
	fmt.Printf("top k frequent is %+v\n", top_k_frequent)

	// top k frequent string max heap
	top_k_frequent_min_heap := heaps.FindKMostFrequentString(arr_str, 2)
	fmt.Printf("top k frequent is with min heap %+v\n", top_k_frequent_min_heap)

	l1 := &models.ListNode{1, &models.ListNode{4, &models.ListNode{5, nil}}}
	l2 := &models.ListNode{1, &models.ListNode{3, &models.ListNode{4, nil}}}
	l3 := &models.ListNode{2, &models.ListNode{6, nil}}

	merged := heaps.ComibineSortedLinkedList([]*models.ListNode{l1, l2, l3})

	for merged != nil {
		print(merged.Val, " ")
		merged = merged.Next
	}

	m := heaps.NewMedianFinder()
	stream := []int{10, 2, 5, 7, 3, 9, 1, 6, 8, 4}
	for _, x := range stream {
		m.Add(x)
		fmt.Printf("Added %-2d -> median = %.1f\n", x, m.FindMedian())
	}

	input := [][]int{
		{1, 3},
		{2, 6},
		{8, 10},
		{15, 18},
	}
	fmt.Println("Merged:", intervals.MergeOverlappingIntervals(input))

	// find overlap intervals
	interval1 := [][]int{
		{1, 4},
		{5, 6},
		{9, 10},
	}
	interval2 := [][]int{
		{2, 7},
		{8, 9},
	}
	fmt.Println("Overlap:", intervals.FindAllIntervalsOverlap(interval1, interval2))

	overlaps := [][]int{
		{1, 3},
		{2, 6},
		{4, 8},
		{6, 7},
		{5, 7},
	}
	overlap_count := intervals.FindLargestIntervalsOverlap(overlaps)
	fmt.Println("Overlap count:", overlap_count)

	/// sums between range // tc: O(n)

	// sc: O(n)
	nums = []int{3, -7, 6, 0, -2, 5}
	prefix_sums := []int{nums[0]}
	for i := 1; i < len(nums); i++ {
		prefix_sums = append(prefix_sums, prefix_sums[i-1]+nums[i])
	}
	sum_prefix := prefixsums.FindSumBetweenRange(2, 5, prefix_sums)
	fmt.Println("Sum between range:", sum_prefix)

	///
	sum_array_k := []int{1, 2, -1, 1, 2}
	k := 3
	total := prefixsums.CalculateKSumSubarray(sum_array_k, k)
	fmt.Println("Total:", total)

	//
	nums = []int{2, 3, 1, 4, 5}
	ans := prefixsums.CalculateProductArrayExceptSelf(nums)
	fmt.Println("Product array except self:", ans)

	// trees
	// invert binary tree
	root := &models.TreeNode{Val: 4}
	root.Left = &models.TreeNode{Val: 2}
	root.Right = &models.TreeNode{Val: 7}
	root.Left.Left = &models.TreeNode{Val: 1}
	root.Left.Right = &models.TreeNode{Val: 3}
	root.Right.Left = &models.TreeNode{Val: 6}
	root.Right.Right = &models.TreeNode{Val: 9}

	tree.InOrderTraversalOfTree(root)

	root_return := tree.InvertBinaryTree(root)

	tree.InOrderTraversalOfTree(root_return)

	root_t := &models.TreeNode{Val: 1}
	root_t.Left = &models.TreeNode{Val: 2}
	root_t.Left.Left = &models.TreeNode{Val: 3}
	root_t.Left.Left.Left = &models.TreeNode{Val: 4}

	is_balanced_tree := tree.BalancedBinaryTreeValidation(root_t)
	fmt.Println(is_balanced_tree)

	root_right := &models.TreeNode{Val: 1}
	root_right.Left = &models.TreeNode{Val: 2}
	root_right.Right = &models.TreeNode{Val: 3}
	root_right.Left.Right = &models.TreeNode{Val: 5}
	root_right.Right.Right = &models.TreeNode{Val: 4}

	fmt.Println(tree.FindRightMostNodeInBinaryTree(root_right)) // Output: [1 3 4]

	freq_arr := []int{3, 2, 3}
	fmt.Println(find_most_frequent.FindTheMostFrequent(freq_arr)) // Output: 2

	// find widest binary tree
	wide_root := &models.TreeNode{Val: 1}
	wide_root.Left = &models.TreeNode{Val: 2}
	wide_root.Right = &models.TreeNode{Val: 3}

	wide_root.Left.Left = &models.TreeNode{Val: 4}
	wide_root.Left.Right = &models.TreeNode{Val: 5}
	wide_root.Right.Right = &models.TreeNode{Val: 7}

	wide_root.Left.Left.Left = &models.TreeNode{Val: 8}
	wide_root.Left.Left.Right = &models.TreeNode{Val: 9}
	wide_root.Left.Right.Right = &models.TreeNode{Val: 11}
	wide_root.Right.Right.Right = &models.TreeNode{Val: 14} // Left of 6 is nil
	fmt.Println(tree.FindWidestBinaryLevel(wide_root))

	bst_root := &models.TreeNode{Val: 5}
	bst_root.Left = &models.TreeNode{Val: 2}
	bst_root.Right = &models.TreeNode{Val: 7}
	bst_root.Left.Left = &models.TreeNode{Val: 1}
	bst_root.Left.Right = &models.TreeNode{Val: 6}
	bst_root.Right.Left = &models.TreeNode{Val: 7}
	bst_root.Right.Right = &models.TreeNode{Val: 9}

	fmt.Println(tree.BinarySearchTreeValidation(bst_root))

	lca_root := &models.TreeNode{Val: 1}
	lca_root.Left = &models.TreeNode{Val: 2}
	lca_root.Right = &models.TreeNode{Val: 3}
	lca_root.Left.Left = &models.TreeNode{Val: 4}
	lca_root.Left.Right = &models.TreeNode{Val: 5}
	lca_root.Right.Left = &models.TreeNode{Val: 6}
	lca_root.Right.Right = &models.TreeNode{Val: 7}
	lca_root.Right.Left.Left = &models.TreeNode{Val: 8}

	node_p := &models.TreeNode{Val: 8}
	node_q := &models.TreeNode{Val: 7}
	lca := tree.LowestCommonAncestor(lca_root, node_p, node_q)
	fmt.Println()
	fmt.Printf("LCA struct: %+v\n", lca.Val)

	preorder := []int{5, 9, 2, 3, 4, 7}
	inorder := []int{2, 9, 5, 4, 3, 7}
	root_pre_in := tree.BuildBinaryTree(preorder, inorder)
	fmt.Printf("Preorder: %+v\n", root_pre_in)
	PrintInorder(root_pre_in)

}

func PrintInorder(root *models.TreeNode) {
	if root == nil {
		return
	}
	PrintInorder(root.Left)
	fmt.Print(root.Val, " ")
	PrintInorder(root.Right)
}
