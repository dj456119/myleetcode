/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2021-12-31 11:36:22
 * @LastEditors: cm.d
 * @LastEditTime: 2021-12-31 11:37:50
 */
package myleetcode

func maxSubArray(nums []int) int {
	max := 0
	sum := 0
	for _, j := range nums {
		if sum > sum+j {
			max = sum
		} else {
			max = sum + j
		}
		sum = sum + j
		if sum < 0 {
			sum = 0
		}
	}
	return max
}
