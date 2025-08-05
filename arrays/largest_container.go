package arrays

// Tc:= O(n)
// Sc:= O(1)
func LargestContainer(nums []int) int {
	max_water := 0
	i, j := 0, len(nums)-1

	for i < j {
		cal_water := (j - i) * min(nums[i], nums[j])
		if cal_water > max_water {
			max_water = cal_water
		}

		if nums[i] < nums[j] {
			i += 1
		} else {
			j -= 1
		}
	}
	return max_water
}
