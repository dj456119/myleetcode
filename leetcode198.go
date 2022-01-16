/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2022-01-16 13:40:11
 * @LastEditors: cm.d
 * @LastEditTime: 2022-01-16 14:16:11
 */
package myleetcode

func rob(nums []int) int {
	arr := make([]int, len(nums))
	result := 0
	for i := range arr {
		if i == 0 {
			arr[i] = nums[i]

		} else {
			if i == 1 {
				arr[i] = max(nums[i], arr[i-1])
			} else {
				arr[i] = max(arr[i-2]+nums[i], arr[i-1])
			}
		}
		if arr[i] > result {
			result = arr[i]
		}
	}
	return result
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
