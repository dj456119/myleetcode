package myleetcode

import "sort"

func maximumGroups(grades []int) int {
	sort.Slice(grades, func(i, j int) bool { return grades[i] < grades[j] })
	z := 1
	f := [][]int{}
	for {
		f = append(f, []int{})
		for i := 0; i < z; i++ {
			f[z-1] = append(f[z-1], grades[0])
			grades = grades[1:]
			if len(grades) == 0 {
				if len(f) == 1 {
					return 1
				}
				if len(f[len(f)-1]) <= len(f[len(f)-2]) {
					return len(f) - 1
				}
				return len(f)
			}
		}
		z++
	}
	return -1
}
