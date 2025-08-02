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
}
