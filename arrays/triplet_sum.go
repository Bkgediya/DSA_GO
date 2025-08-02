package arrays

import (
	"sort"
)

func TripletSum(nums []int) [][]int {
	// sort the array
	sort.Ints(nums)
	var result [][]int

	for i := 0; i < len(nums); i++ {
		// find the pair of two sum
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		pairs := pair_of_two_sum(nums, i+1, -nums[i])

		for _, p := range pairs {
			result = append(result, []int{nums[i], p[0], p[1]})
		}
	}

	return result
}

func pair_of_two_sum(nums []int, start int, target int) [][]int {
	var pairs [][]int

	b := start
	c := len(nums) - 1

	for b < c {
		if nums[b]+nums[c] == target {
			pairs = append(pairs, []int{nums[b], nums[c]})
			b += 1
			if b < c && nums[b] == nums[b-1] {
				b += 1
			}
		} else if nums[b]+nums[c] > target {
			c--
		} else if nums[b]+nums[c] < target {
			b++
		}
	}

	return pairs
}
