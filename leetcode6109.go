package myleetcode

func peopleAwareOfSecret(n int, delay int, forget int) int {
	c := make([]int, n)
	c[0] = 1

	for i := 1; i < n; i++ {
		i1 := 0
		if i-forget > 0 {
			i1 = i - forget
		}
		i2 := 0
		if i-delay > 0 {
			i2 = i - delay
		}
		c[i] = c[i-1] + c[i2] - c[i1]
	}
	return c[len(c)-1]
}
