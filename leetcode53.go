/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2021-12-31 11:36:22
 * @LastEditors: cm.d
 * @LastEditTime: 2022-01-15 23:16:54
 */
package myleetcode

func maxSubArray(nums []int) int {
	max := nums[0]
	lxsum := nums[0]

	for i, j := range nums {
		if i == 0 {
			continue
		}
		lxsum = lxsum + j
		if lxsum < j {
			lxsum = j
		}
		if lxsum > max {
			max = lxsum
		}
	}
	return max
}
