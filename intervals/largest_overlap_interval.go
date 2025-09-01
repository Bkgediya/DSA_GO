package intervals

import "sort"

type Point struct {
	val int
	typ rune // 's' = start, 'e' = end
}

func FindLargestIntervalsOverlap(intervals [][]int) int {
	points := []Point{}
	interval_largest := 0
	for _, interval := range intervals {
		points = append(points, Point{interval[0], 's'})
		points = append(points, Point{interval[1], 'e'})
	}

	sort.Slice(points, func(i, j int) bool {
		return points[i].val < points[j].val
	})
	active_internals := 0

	for _, point := range points {
		if point.typ == 's' {
			active_internals++
		} else {
			active_internals--
		}

		interval_largest = max(interval_largest, active_internals)
	}

	return interval_largest
}
