package myleetcode

func findWinners(matches [][]int) [][]int {
	max := -1
	z2 := make(map[int]int)
	for i := range matches {
		if matches[i][0] > max {
			max = matches[i][0]
		}
		if matches[i][1] > max {
			max = matches[i][1]
		}
	}
	z1 := make([]int, max+1)
	for i := range matches {
		z1[matches[i][1]]++
		z2[matches[i][0]]++
		z2[matches[i][1]]++
	}
	result1 := []int{}
	result2 := []int{}
	for i := range z1 {
		if _, ok := z2[i]; ok {
			if z1[i] == 1 {
				result2 = append(result2, i)
			}
			if z1[i] == 0 {
				result1 = append(result1, i)
			}
		}
	}
	return [][]int{result1, result2}
}
