package graph

// Tc : O(M.N)
// Sc : O(M.N)
func MatrixInfection(nums [][]int) int {
	dirs := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	ones, seconds := 0, 0
	queue := make([][]int, 0)
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums[0]); j++ {
			if nums[i][j] == 1 {
				ones++
			} else if nums[i][j] == 2 {
				queue = append(queue, []int{i, j})
			}
		}
	}

	for len(queue) > 0 && ones > 0 {
		seconds += 1
		for i := 0; i < len(queue); i++ {
			row, col := queue[0][0], queue[0][1]
			queue = queue[1:]
			for _, dir := range dirs {
				next_row, next_col := row+dir[0], col+dir[1]
				if is_within_bound(nums, next_row, next_col) && nums[next_row][next_col] == 1 {
					nums[next_row][next_col] = 2
					ones -= 1
					queue = append(queue, []int{next_row, next_col})
				}
			}
		}

	}

	if ones > 0 {
		return -1
	}
	return seconds
}
