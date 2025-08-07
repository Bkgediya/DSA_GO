package main

import (
	"DSA_Go/arrays"
	hashmapsets "DSA_Go/hashmap_sets"
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
	fmt.Printf("%+v", zero_stripped_arr)

}
