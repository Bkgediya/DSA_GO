package arrays

// TC : O(n)
// SC : O(1)
func TwoSum(nums []int, target int) []int {
	i := 0
	j := len(nums) - 1

	for i < j {
		if nums[i]+nums[j] == target {
			return []int{i, j}
		} else if nums[i]+nums[j] > target {
			j--
		} else {
			i++
		}
	}
	return []int{}
}
