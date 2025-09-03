package prefixsums

// tc: O(n)
// sc: O(n)
func CalculateKSumSubarray(nums []int, k int) int {
	total := 0
	sum_freq_map := make(map[int]int)
	sum_freq_map[0] = 1

	curr_prefix_sum := 0

	for num := range nums {
		curr_prefix_sum += num
		if _, exist := sum_freq_map[curr_prefix_sum-k]; exist {
			total += sum_freq_map[curr_prefix_sum-k]
		}
		sum_freq_map[curr_prefix_sum] += 1
	}

	return total
}
