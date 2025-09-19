package graph

// tc : O(M.N)
// sc :O(M.N)
func FindIsland(nums [][]int) int {
	count := 0
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums[0]); j++ {
			if nums[i][j] == 1 {
				dfsIsland(nums, i, j)
				count++
			}

		}
	}
	return count
}

func dfsIsland(nums [][]int, row, col int) {
	nums[row][col] = -1

	dirs := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	for _, dir := range dirs {
		next_row, next_col := row+dir[0], col+dir[1]

		if is_within_bound(nums, next_row, next_col) && nums[next_row][next_col] == 1 {
			dfsIsland(nums, next_row, next_col)
		}
	}
}

func is_within_bound(nums [][]int, row, col int) bool {
	return row >= 0 && row < len(nums) && col >= 0 && col < len(nums[0])
}
