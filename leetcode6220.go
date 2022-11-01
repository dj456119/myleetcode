package myleetcode

func averageValue(nums []int) int {
	sum := 0
	count := 0
	for i := range nums {
		if nums[i]%3 == 0 && nums[i]%2 == 0 {
			sum = sum + nums[i]
			count++
		}
	}
	if count == 0 {
		return 0
	}
	return sum / count
}
