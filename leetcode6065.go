package myleetcode

func largestCombination(candidates []int) int {
	resultArr := make([]int, 32)
	for i := 0; i < 32; i++ {
		for _, num := range candidates {
			z := num >> i & (1)
			if z == 1 {
				resultArr[i]++
			}
		}
	}
	max := 0
	for i := range resultArr {
		if resultArr[i] > max {
			max = resultArr[i]
		}
	}
	return max
}
