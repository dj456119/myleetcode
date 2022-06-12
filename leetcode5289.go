package myleetcode

func distributeCookies(cookies []int, k int) int {
	return backtrace5289(cookies, 0, make([]int, k))
}

func backtrace5289(cookies []int, now int, sums []int) int {
	if now == len(cookies) {
		max := 0
		for i := range sums {
			if sums[i] > max {
				max = sums[i]
			}
		}
		return max
	}
	min := 9999999999
	for i := 0; i < len(sums); i++ {
		bc := make([]int, len(sums))
		copy(bc, sums)
		bc[i] = bc[i] + cookies[i]
		m := backtrace5289(cookies, now+1, bc)
		if m < min {
			min = m
		}
	}
	return min
}
