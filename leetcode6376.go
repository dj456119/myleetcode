package myleetcode

func rowAndMaximumOnes(mat [][]int) []int {
	result := []int{}
	max := 0
	for i := range mat {
		count := 0
		for j := range mat[i] {
			if mat[i][j] == 1 {
				count++
			}
		}
		if count > max {
			max = count
			result = []int{i, max}
		}
	}
	return result
}
