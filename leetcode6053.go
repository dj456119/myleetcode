package myleetcode

func countUnguarded(m int, n int, guards [][]int, walls [][]int) int {
	mm := make([][]int, m)
	for i := range mm {
		mm[i] = make([]int, n)
	}
	for i := range guards {
		mm[guards[i][0]][guards[i][1]] = 1
	}
	for i := range walls {
		mm[walls[i][0]][walls[i][1]] = 2
	}

	count := 0
	heng := make([]int, n)
	shu := make([]int, m)
	for y := range mm {
		for x := range mm[y] {
			if mm[y][x] == 2 {
				heng[x] = 2
				shu[y] = 2
			}
			if mm[y][x] == 1 {
				heng[x] = 1
				shu[y] = 1
			}
			if mm[y][x] == 0 {
				if heng[x] == 1 || shu[y] == 1 {
					mm[y][x] = 3
				}
			}
		}
	}
	heng = make([]int, n)
	shu = make([]int, m)
	for y := m - 1; y >= 0; y-- {
		for x := n - 1; x >= 0; x-- {
			if mm[y][x] == 2 {
				heng[x] = 2
				shu[y] = 2
			}
			if mm[y][x] == 1 {
				heng[x] = 1
				shu[y] = 1
			}
			if mm[y][x] == 0 {
				if heng[x] == 1 || shu[y] == 1 {
					mm[y][x] = 3
				}
			}
		}
	}

	for y := range mm {
		for x := range mm[y] {
			if mm[y][x] == 0 {
				count++
			}
		}
	}
	return count
}
