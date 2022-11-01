package myleetcode

func findMaxK(nums []int) int {
	result := -1
	for i := range nums {
		if nums[i] < 0 || nums[i] < result {
			continue
		}
		for j := range nums {
			if nums[i] == -nums[j] && nums[i] > result {
				result = nums[i]
			}
		}
	}
	return result
}
