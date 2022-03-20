package myleetcode

func countHillValley(nums []int) int {
	result := 0

	for i := 0; i < len(nums); i++ {
		if i == 0 || i == len(nums)-1 {
			continue
		}
		if nums[i] < nums[i-1] && nums[i] < nums[i+1] {
			result++
			continue
		}
		if nums[i] > nums[i-1] && nums[i] > nums[i+1] {
			result++
			continue
		}
		if nums[i] == nums[i+1] {
			nums[i] = nums[i-1]
		}
	}
	return result
}
