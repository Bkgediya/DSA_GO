package stack

// tc: O(n)
// sc: O(n)
func FindNextLargeNumberToRight(nums []int) []int {
	n := len(nums)
	result := make([]int, n)

	stack := []int{}
 
	for i := n - 1; i >= 0; i-- {
		for len(stack) > 0 && stack[len(stack)-1] < nums[i] {
			stack = stack[:len(stack)-1]	
		}
		if len(stack) == 0 {
			result[i] = -1
		} else {
			result[i] = stack[len(stack)-1]
		}
		stack = append(stack, nums[i])
	}
	return result
}
