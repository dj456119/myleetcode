package myleetcode

import "sort"

func dividePlayers(skill []int) int64 {
	sort.Slice(skill, func(i, j int) bool { return skill[i] < skill[j] })
	z := skill[0] + skill[len(skill)-1]
	start := 0
	end := len(skill) - 1
	m := [][]int{}
	for start < end {
		c := skill[start] + skill[end]
		if c != z {
			return -1
		}
		m = append(m, []int{skill[start], skill[end]})
		start++
		end--
	}
	var result int64
	for i := range m {
		result = result + int64(m[i][0])*int64(m[i][1])
	}
	return result
}
