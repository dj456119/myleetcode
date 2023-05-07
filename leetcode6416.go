package myleetcode

func distinctDifferenceArray(nums []int) []int {
	n := len(nums)
	diff := make([]int, n)

	prefixCount := make(map[int]int)
	suffixCount := make(map[int]int)

	for _, num := range nums {
		suffixCount[num]++
	}

	for i, num := range nums {
		suffixCount[num]--
		if suffixCount[num] == 0 {
			delete(suffixCount, num)
		}
		prefixCount[num]++

		diff[i] = len(prefixCount) - len(suffixCount)
	}

	return diff
}
