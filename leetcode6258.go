package myleetcode

func longestSquareStreak(nums []int) int {
	z := make(map[int]int)
	result := 0
	for i := range nums {
		z[nums[i]] = 1
	}
	for i := range nums {
		c := nums[i] * nums[i]
		cc := 1
		for {
			if _, ok := z[c]; !ok {
				break
			}
			cc++
			c = c * c
		}
		if cc > result {
			result = cc
		}
	}
	if result == 1 {
		return -1
	}
	return result
}
