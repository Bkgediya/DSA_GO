package prefixsums

func FindSumBetweenRange(start, end int, prefix_sums []int) int {
	if start == 0 {
		return prefix_sums[end]
	}

	total := prefix_sums[end] - prefix_sums[start-1]
	return total
}
