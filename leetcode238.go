/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2022-01-31 19:17:07
 * @LastEditors: cm.d
 * @LastEditTime: 2022-01-31 19:21:27
 */
package myleetcode

func productExceptSelf(nums []int) []int {
	result := []int{}
	ji := 0
	for i, j := range nums {
		if i == 0 {
			ji = j
		} else {
			ji = ji * j
		}
		result = append(result, ji)
	}
	right := 1
	for i := len(nums) - 1; i >= 0; i-- {
		if i == 0 {
			result[i] = right
		}
		result[i] = result[i-1] * right
		right = right * nums[i]
	}
	return result
}
