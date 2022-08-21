package myleetcode

func largestLocal(grid [][]int) [][]int {
	length := len(grid) - 2
	result := make([][]int, length)
	for i := range result {
		result[i] = make([]int, length)
	}
	for y := range grid {
		if y == length {
			break
		}
		for x := range grid[y] {
			if x == length {
				break
			}
			max := 0
			for j := y; j < y+3; j++ {
				for i := x; i < x+3; i++ {
					if grid[j][i] > max {
						max = grid[j][i]
					}
				}
			}
			result[y][x] = max
		}
	}
	return result
}
