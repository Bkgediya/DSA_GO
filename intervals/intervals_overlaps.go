package intervals

// time complexity : O(n+m) | space complexity : O(n)
func FindAllIntervalsOverlap(interval1 [][]int, interval2 [][]int) [][]int {
	result := [][]int{}

	i, j := 0, 0
	a, b := []int{}, []int{}
	for i < len(interval1) && j < len(interval2) {
		if interval1[i][0] <= interval2[j][0] {
			a, b = interval1[i], interval2[j]
		} else {
			a, b = interval2[j], interval1[i]
		}
		if a[1] >= b[0] {
			result = append(result, []int{b[0], min(a[1], b[1])})
		}

		if interval1[i][1] <= interval2[j][1] {
			i++
		} else {
			j++
		}
	}

	return result
}
