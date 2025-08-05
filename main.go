package main

import (
	"DSA_Go/arrays"
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

	nums_container := []int{2, 7, 8, 3, 7, 6}
	max_water := arrays.LargestContainer(nums_container)
	fmt.Printf("%+v", max_water)
}
