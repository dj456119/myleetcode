package myleetcode

import "fmt"

func findColumnWidth(grid [][]int) []int {
	result := make([]int, len(grid[0]))
	for i := range grid {
		for j := range grid[i] {
			q := fmt.Sprintf("%d", grid[i][j])
			if len(q) > result[j] {
				result[j] = len(q)
			}
		}
	}
	return result
}
