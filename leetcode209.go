package myleetcode

func minSubArrayLen(target int, nums []int) int {
	if len(nums) == 1 {
		if nums[0] >= target {
			return 1
		}
		return 0
	}
	sum := 0
	start := 0
	result := 99999999
	for i := range nums {
		nowl := i - start + 1
		if sum+nums[i] >= target {
			if nowl < result {
				result = nowl
			}
			sum = sum + nums[i]
			for j := start; j <= i; j++ {
				sum = sum - nums[j]
				if sum < target {
					start = j + 1
					break
				} else {
					nowl = i - j
					if nowl < result {
						result = nowl
					}
				}
			}
		} else {
			sum = sum + nums[i]
		}
	}
	if result == 99999999 {
		return 0
	}
	return result
}
