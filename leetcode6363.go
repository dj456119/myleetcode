package myleetcode

import "sort"

func findMatrix(nums []int) [][]int {
	q := nums
	sort.Slice(q, func(i, j int) bool { return q[i] < q[j] })
	result := [][]int{}
	pos := 0
	for i := range q {
		if i == 0 {
			z := []int{q[i]}
			result = append(result, z)
			continue
		}
		if q[i] == q[i-1] {
			pos++
			if len(result) <= pos {
				z := []int{q[i]}
				result = append(result, z)
			} else {
				result[pos] = append(result[pos], q[i])
			}
		} else {
			pos = 0
			result[pos] = append(result[pos], q[i])
		}
	}
	return result
}
