package myleetcode

func applyOperations(nums []int) []int {
	for i := range nums {
		if i == len(nums)-1 {
			break
		}
		if nums[i] == nums[i+1] {
			nums[i] = 2 * nums[i]
			nums[i+1] = 0
		}
	}
	count := 0
	result := []int{}
	for i := range nums {
		if nums[i] != 0 {
			count++
			result = append(result, nums[i])
		}
	}
	for i := 0; i < len(nums)-count; i++ {
		result = append(result, 0)
	}
	return result
}
