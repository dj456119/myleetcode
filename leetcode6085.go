package myleetcode

import "sort"

func maximumImportance(n int, roads [][]int) int64 {
	c := make([]int, n)
	for i := range roads {
		c[roads[i][0]]++
		c[roads[i][1]]++
	}
	z := make([][]int, n)
	for i := range z {
		z[i] = make([]int, 2)
		z[i][0] = i
		z[i][1] = c[i]
	}
	sort.Slice(z, func(i, j int) bool {
		return z[i][1] < z[j][1]
	})
	f := make([]int, n)
	for i := range z {
		f[z[i][0]] = i + 1
	}
	var sum int64
	for i := range roads {
		sum = sum + int64(f[roads[i][0]]) + int64(f[roads[i][1]])
	}
	return sum
}
