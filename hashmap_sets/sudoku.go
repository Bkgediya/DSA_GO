package hashmapsets

func ValidateSudoku(nums [][]int) bool {
	// struct{} is anonymous struct type with no field
	rows := make(map[int]struct{})
	cols := make(map[int]struct{})
	grids := make(map[int]struct{})

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			num := nums[i][j]

			if num == 0 {
				continue
			}

			if _, ok := rows[num]; ok {
				return false
			}

			if _, ok := cols[num]; ok {
				return false
			}
			if _, ok := grids[num]; ok {
				return false
			}

			// invoking struct type
			rows[num] = struct{}{}
			cols[num] = struct{}{}
			grids[num] = struct{}{}
		}
	}

	return true

}
