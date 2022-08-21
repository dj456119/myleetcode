package myleetcode

func longestCycle(edges []int) int {
	mm := make([]int, len(edges))
	result := 0
	for i := range edges {
		cc := make([]int, len(edges))
		n := i
		mm[i] = 1
		l1 := 1
		cc[n] = l1
		n = edges[n]
		for {
			if cc[n] != 0 {
				m := l1 + 1 - cc[n]
				if m > result {
					result = m
				}
			}
			if cc[n] == -1 {
				break
			}
			if mm[i] == 1 {
				break
			}
			l1
			cc[n] = l1
			n = cc[n]
		}
	}
	return result
}
