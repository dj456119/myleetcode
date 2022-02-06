/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2022-02-06 17:34:39
 * @LastEditors: cm.d
 * @LastEditTime: 2022-02-06 17:41:13
 */
package myleetcode

func findUnsortedSubarray(nums []int) int {
	max := nums[0]
	right := 0
	for i := range nums {
		if nums[i] < max {
			right = i
		} else {
			max = nums[i]
		}
	}
	min := nums[len(nums)-1]
	left := len(nums) - 1
	for i := left; i >= 0; i-- {
		if nums[i] > min {
			left = i
		} else {
			min = nums[i]
		}
	}
	if right == 0 && left == len(nums)-1 {
		return 0
	}
	return right - left + 1
}
