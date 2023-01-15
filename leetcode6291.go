package myleetcode

import "math"

func differenceOfSum(nums []int) int {
	result1 := 0
	result2 := 0
	for i := range nums {
		result1 = nums[i] + result1
		j := nums[i]
		for j != 0 {
			z := j % 10
			j = j / 10
			result2 = result2 + z
		}
	}
	return int(math.Abs(float64(result1) - float64(result2)))
}
