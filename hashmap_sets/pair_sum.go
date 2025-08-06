package hashmapsets

//Tc:= O(n)
//Sc:= O(n)
func PairSum(nums []int, target int) []int {
	m := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		val := target - nums[i]
		if _, ok := m[val]; ok {
			return []int{m[val], i}
		} else {
			m[nums[i]] = i
		}
	}

	return []int{}
}
