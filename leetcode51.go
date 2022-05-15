package myleetcode

func solveNQueens(n int) [][]string {
	return backtrace51(n, 0, []string{}, make(map[int]bool), make(map[int]bool), make(map[int]bool))
}

func backtrace51(n int, pos int, result []string, left map[int]bool, right map[int]bool, zong map[int]bool) [][]string {
	rrr := [][]string{}
	if pos == n {
		return append(rrr, result)
	}
	for i := 0; i < n; i++ {
		if isOK(pos, i, left, right, zong, n) {
			c := make([]string, len(result))
			copy(c, result)
			now := []byte{}
			for j := 0; j < n; j++ {
				if j == i {
					now = append(now, 'Q')
				} else {
					now = append(now, '.')
				}
			}
			c = append(c, string(now))
			rr := backtrace51(n, pos+1, c, left, right, zong)
			if rr != nil {
				rrr = append(rrr, rr...)
			}
			delete(left, i-pos)
			delete(right, (n-i)-pos)
			delete(zong, i)
		}
	}
	return rrr
}

func isOK(pos, i int, left map[int]bool, right map[int]bool, zong map[int]bool, n int) bool {
	if left[i-pos] {
		return false
	}
	if right[(n-i)-pos] {
		return false
	}
	if zong[i] {
		return false
	}
	left[i-pos] = true
	right[(n-i)-pos] = true
	zong[i] = true
	return true
}
