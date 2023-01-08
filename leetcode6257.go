package myleetcode

func deleteGreatestValue(grid [][]int) int {
	result := 0

	for len(grid[0]) != 0 {
		m2 := 0
		gt := [][]int{}
		for i := range grid {
			max := 0
			for j := range grid[i] {
				if grid[i][j] > max {
					max = grid[i][j]
				}
			}
			if max > m2 {
				m2 = max
			}
			z := []int{}
			flag := false
			for j := range grid[i] {
				if grid[i][j] == max && !flag {
					flag = true
					continue
				}
				z = append(z, grid[i][j])
			}
			gt = append(gt, z)
		}
		grid = gt
		result = result + m2
	}
	return result
}
