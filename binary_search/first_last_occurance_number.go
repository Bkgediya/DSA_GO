package binarysearch

func FirstLastOccurance(nums []int, target int) []int {
	lowerBound := lowerBound(nums, target)
	upperBound := upperBound(nums, target)
	return []int{lowerBound, upperBound}
}

func lowerBound(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := (left + right) / 2
		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			right = mid
		}
	}

	if nums[left] == target {
		return left
	} else {
		return -1
	}

}

func upperBound(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := ((left + right) / 2) + 1
		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid
		}
	}

	if nums[right] == target {
		return right
	} else {
		return -1
	}
}
