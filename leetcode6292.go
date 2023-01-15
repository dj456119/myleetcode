package myleetcode

func rangeAddQueries(n int, queries [][]int) [][]int {
	q := make([][]int, n)
	for i := range q {
		q[i] = make([]int, n)
	}
	xx := make([]int, n)
	yy := make([]int, n)
	for i := range queries {
		s0 := queries[i][0]
		e0 := queries[i][1]
		s1 := queries[i][2]
		e1 := queries[i][3]
		for i := s0; i <= s1; i++ {
			xx[i]++
		}
		for j := e0; j <= e1; j++ {
			yy[j]++
		}
	}
	for y := range q {
		for x := range q[y] {
			z := yy[y]
			if xx[x] < yy[y] {
				z = xx[x]
			}
			q[x][y] = z
		}
	}
	return q
}
