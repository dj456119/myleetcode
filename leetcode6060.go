package myleetcode

func findClosestNumber(nums []int) int {
	max := 99999999
	for i := range nums {
		if nums[i] < 0 {
			if -nums[i] < max {
				max = -nums[i]
			}
		}
		if nums[i] == 0 {
			return 0
		}
		if nums[i] > 0 {
			if nums[i] < max {
				max = nums[i]
			}
		}
	}
	for i := range nums {
		if nums[i] == max {
			return max
		}
	}
	return -max
}
