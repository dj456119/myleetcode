package myleetcode

func findPrefixScore(nums []int) []int64 {
	result := []int64{}
	max := 0
	for i := range nums {
		if nums[i] > max {
			max = nums[i]
		}
		result = append(result, int64(max+nums[i]))
	}
	var sum int64
	for i := range result {
		sum = result[i] + sum
	}
	result1 := make([]int64, len(result))
	var right int64
	for i := len(result) - 1; i >= 0; i-- {
		result1[i] = sum - right
		right = right + result[i]
	}
	return result1
}
