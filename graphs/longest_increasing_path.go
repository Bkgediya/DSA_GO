package graph

// tc : O(M*N)
// sc : O(M*N)
func LongestIncreasingPath(matrix [][]int) int {
	if len(matrix) == 0 {
		return 0
	}

	res := 0
	memo := make([][]int, len(matrix))
	for i := range memo {
		memo[i] = make([]int, len(matrix[0]))
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			res = max(res, DfSLongestIncreasingPath(matrix, i, j, memo))
		}
	}

	return res
}

func DfSLongestIncreasingPath(matrix [][]int, row, col int, memo [][]int) int {
	if memo[row][col] != 0 {
		return memo[row][col]
	}

	max_path := 1
	dirs := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	for _, dir := range dirs {
		next_row := row + dir[0]
		next_col := col + dir[1]

		if is_within_bound(matrix, next_row, next_col) && matrix[next_row][next_col] > matrix[row][col] {
			max_path = max(max_path, 1+DfSLongestIncreasingPath(matrix, next_row, next_col, memo))
		}
	}

	memo[row][col] = max_path
	return max_path
}
