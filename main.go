package main

import (
	"DSA_Go/arrays"
	binarysearch "DSA_Go/binary_search"
	hashmapsets "DSA_Go/hashmap_sets"
	linkedlist "DSA_Go/linked_list"
	"DSA_Go/linked_list/models"
	"DSA_Go/sliding_window"
	"DSA_Go/stack"
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
}
