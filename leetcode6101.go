package myleetcode

func checkXMatrix(grid [][]int) bool {
	for y := range grid {
		for x := range grid[y] {
			if x != y && y+x != len(grid)-1 {
				continue
			}
			if grid[y][x] == 0 {
				return false
			}
		}
	}

	for y := range grid {
		for x := range grid[y] {
			if x == y || y+x == len(grid)-1 {
				continue
			}
			if grid[y][x] != 0 {
				return false
			}
		}
	}
	return true
}
