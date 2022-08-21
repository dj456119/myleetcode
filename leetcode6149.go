package myleetcode

func edgeScore(edges []int) int {
	c := make([][]int, len(edges))
	for i := range edges {
		c[edges[i]] = append(c[edges[i]], i)
	}
	mS := 0
	result := 0
	for i := range c {
		m := 0
		for j := range c[i] {
			m = m + c[i][j]
		}
		if m < mS {
			continue
		}
		if m == mS {
			continue
		}
		mS = m
		result = i
	}
	return result
}
