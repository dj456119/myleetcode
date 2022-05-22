package myleetcode

import "sort"

func maximumBags(capacity []int, rocks []int, additionalRocks int) int {
	result := 0
	c := make([]int, len(capacity))
	for i := range capacity {
		c[i] = capacity[i] - rocks[i]
	}
	sort.Slice(c, func(i, j int) bool { return c[i] < c[j] })
	for i := range capacity {
		if c[i] == 0 {
			continue
		}
		if additionalRocks >= c[i] {
			additionalRocks = additionalRocks - c[i]
			c[i] = 0
		}
	}
	for i := range capacity {
		if c[i] == 0 {
			result++
		}
	}
	return result
}
