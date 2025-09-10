package hashmapsets

// time complexity O(n^2)
// space complexity O(n^2)
func ValidateSudoku(nums [][]int) bool {
	// struct{} is anonymous struct type with no field
	rows := make([]map[int]struct{}, 9)
	cols := make([]map[int]struct{}, 9)
	grids := make([]map[int]struct{}, 9)

	for i := 0; i < 9; i++ {
		rows[i] = make(map[int]struct{})
		cols[i] = make(map[int]struct{})
		grids[i] = make(map[int]struct{})
	}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			num := nums[i][j]

			if num == 0 {
				continue
			}

			if _, ok := rows[i][num]; ok {
				return false
			}

			if _, ok := cols[j][num]; ok {
				return false
			}

			// grid index
			gridIndex := (i/3)*3 + (j / 3)
			if _, ok := grids[gridIndex][num]; ok {
				return false
			}

			// invoking struct type
			rows[i][num] = struct{}{}
			cols[j][num] = struct{}{}
			grids[gridIndex][num] = struct{}{}
		}
	}

	return true

}
