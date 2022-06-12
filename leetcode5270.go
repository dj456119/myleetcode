package myleetcode

func minPathCost(grid [][]int, moveCost [][]int) int {
	for i := len(grid) - 2; i >= 0; i-- {
		for j := range grid[i] {
			min := 9999999999
			for z := 0; z < len(grid[i+1]); z++ {
				cost := grid[i][j] + grid[i+1][z] + moveCost[grid[i][j]][z]
				if cost < min {
					min = cost
				}
			}
			grid[i][j] = min
		}
	}
	result := 9099999999
	for i := 0; i < len(grid[0]); i++ {
		if grid[0][i] < result {
			result = grid[0][i]
		}
	}
	return result
}
