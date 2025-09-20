package graph

func BiparatiteValidation(nums [][]int) bool {
	n := len(nums)
	color := make([]int, n)
	for i := 0; i < n; i++ {
		if color[i] == 0 && !biparateDFS(i, 1, color, nums) {
			return false
		}
	}

	return true
}

func biparateDFS(node int, c int, color []int, nums [][]int) bool {
	color[node] = c
	for _, neighbour := range nums[node] {
		if color[neighbour] == c {
			return false
		}
		if color[neighbour] == 0 {
			if !biparateDFS(neighbour, -c, color, nums) {
				return false
			}
		}
	}
	return true
}
