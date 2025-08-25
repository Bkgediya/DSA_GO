package binarysearch

func CuttingWood(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left < right {
		mid := ((left + right) / 2) + 1
		if cuts_enough_wood(mid, target, nums) {
			left = mid
		} else {
			right = mid - 1
		}
	}
	return right

}

func cuts_enough_wood(length int, k int, nums []int) bool {
	count := 0
	for _, num := range nums {
		if num > length {
			count += (num - length)
		}
	}
	return count >= k
}
