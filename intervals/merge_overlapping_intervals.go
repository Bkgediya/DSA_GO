package intervals

import "sort"

func MergeOverlappingIntervals(intervalsArr [][]int) [][]int {
	if len(intervalsArr) == 0 {
		return [][]int{}
	}

	sort.Slice(intervalsArr, func(i, j int) bool {
		return intervalsArr[i][0] < intervalsArr[j][0]
	})

	// merge
	result := [][]int{intervalsArr[0]}

	for _, interval := range intervalsArr {
		last := result[len(result)-1]

		if interval[0] <= last[1] {
			if interval[1] > last[1] {
				last[1] = interval[1]
			}
		} else {
			// no overlap
			result = append(result, interval)
		}
	}
	return result
}
