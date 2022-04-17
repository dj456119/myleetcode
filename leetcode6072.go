package myleetcode

func maxTrailingZeros(grid [][]int) int {
	zeroMap := make([][]int, len(grid))
	heng := make([][]int, len(grid))
	shu := make([][]int, len(grid))
	heng2 := make([][]int, len(grid))
	shu2 := make([][]int, len(grid))
	for i := range zeroMap {
		zeroMap[i] = make([]int, len(grid[0]))
	}
	for i := range heng {
		heng[i] = make([]int, len(grid[0]))
	}
	for i := range shu {
		shu[i] = make([]int, len(grid[0]))
	}
	for i := range heng2 {
		heng2[i] = make([]int, len(grid[0]))
	}
	for i := range shu2 {
		shu2[i] = make([]int, len(grid[0]))
	}
	for y := range grid {
		c := 1
		for x := range grid[y] {
			c = c * grid[y][x]
			heng[y][x] = c / grid[y][x]
		}
	}
	for y := range grid {
		c := 1
		for x := len(grid[y]) - 1; x >= 0; x-- {
			c = c * grid[y][x]
			heng2[y][x] = c / grid[y][x]
		}
	}
	for x := 0; x < len(grid[0]); x++ {
		c := 1
		for y := 0; y < len(grid); y++ {
			c = c * grid[y][x]
			shu[y][x] = c / grid[y][x]
		}
	}
	for x := 0; x < len(grid[0]); x++ {
		c := 1
		for y := len(grid) - 1; y >= 0; y-- {
			c = c * grid[y][x]
			shu2[y][x] = c / grid[y][x]
		}
	}
	result := -1
	for y := range heng {
		for x := range heng[y] {
			h1 := heng[y][x] * shu[y][x] * grid[y][x]
			h1_zero := havezero(h1)
			h2 := heng[y][x] * shu2[y][x] * grid[y][x]
			h2_zero := havezero(h2)
			h3 := heng2[y][x] * shu[y][x] * grid[y][x]
			h3_zero := havezero(h3)
			h4 := heng2[y][x] * shu2[y][x] * grid[y][x]
			h4_zero := havezero(h4)
			h_zero := max6072(h1_zero, max6072(h2_zero, max6072(h3_zero, h4_zero)))
			result = max6072(result, h_zero)
		}
	}
	return result
}

func havezero(c int) int {
	result := 0
	for c%10 == 0 {
		result++
		c = c / 10
	}
	return result
}

func max6072(a, b int) int {
	if a > b {
		return a
	}
	return b
}
