package myleetcode

func checkValidGrid(grid [][]int) bool {
	z := make([][]int, len(grid)*len(grid))
	for y := range grid {
		for x := range grid {
			z[grid[y][x]] = []int{x, y}
		}
	}
	if z[0][0] != 0 || z[0][1] != 0 {
		return false
	}
	for i := range z {
		if i == len(z)-1 {
			break
		}
		if z[i][0]-2 == z[i+1][0] && z[i][1]-1 == z[i+1][1] {
			continue
		}
		if z[i][0]-1 == z[i+1][0] && z[i][1]-2 == z[i+1][1] {
			continue
		}
		if z[i][0]+2 == z[i+1][0] && z[i][1]-1 == z[i+1][1] {
			continue
		}
		if z[i][0]+1 == z[i+1][0] && z[i][1]-2 == z[i+1][1] {
			continue
		}
		if z[i][0]-2 == z[i+1][0] && z[i][1]+1 == z[i+1][1] {
			continue
		}
		if z[i][0]-1 == z[i+1][0] && z[i][1]+2 == z[i+1][1] {
			continue
		}
		if z[i][0]+2 == z[i+1][0] && z[i][1]+1 == z[i+1][1] {
			continue
		}
		if z[i][0]+1 == z[i+1][0] && z[i][1]+2 == z[i+1][1] {
			continue
		}
		return false
	}
	return true
}
