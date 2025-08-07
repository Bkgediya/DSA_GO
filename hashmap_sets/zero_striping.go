package hashmapsets

// question: For each zero in an m x n matrix. set its entire row and column to zero in place. 

func ZeroStriping(grid [][]int) [][]int {
	is_first_row_has_zero := false
	is_first_col_has_zero := false

	if len(grid) == 0 || len(grid[0]) == 0 {
		return [][]int{}
	}

	m, n := len(grid), len(grid[0])
	for j := range n {
		if grid[0][j] == 0 {
			is_first_row_has_zero = true
			break
		}
	}

	for i := range m {
		if grid[i][0] == 0 {
			is_first_col_has_zero = true
			break
		}
	}

	for r := range m {
		for c := range n {
			if grid[r][c] == 0 {
				grid[0][c] = 0
				grid[r][0] = 0
			}
		}
	}

	// iterate and based on first row and column make entire row and column zero
	for r := range m {
		for c := range n {
			if grid[r][0] == 0 || grid[0][c] == 0 {
				grid[r][c] = 0
			}
		}
	}

	if is_first_row_has_zero {
		for i := range n {
			grid[0][i] = 0
		}
	}

	if is_first_col_has_zero {
		for i := range m {
			grid[i][0] = 0
		}
	}

	return grid
}
