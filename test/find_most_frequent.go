package find_most_frequent

func FindTheMostFrequent(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	hashmap := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		if _, ok := hashmap[nums[i]]; ok {
			hashmap[nums[i]] += 1
		} else {
			hashmap[nums[i]] = 1
		}

		if hashmap[nums[i]] > len(nums)/2 {
			return nums[i]
		}
	}

	return 0
}
