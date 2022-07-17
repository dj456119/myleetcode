package myleetcode

func numberOfPairs(nums []int) []int {
	cm := make(map[int]int)
	for i := range nums {
		cm[nums[i]]++
	}
	result1 := 0
	result2 := 0
	for _, v := range cm {
		result1 = result1 + v/2
		result2 = result2 + v%2
	}
	return []int{result1, result2}
}
