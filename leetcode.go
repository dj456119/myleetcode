/*
 * @Descripttion:给定一个包含 [0, n] 中 n 个数的数组 nums ，找出 [0, n] 这个范围内没有出现在数组中的那个数。


 * @version:
 * @Author: cm.d
 * @Date: 2021-11-06 00:53:22
 * @LastEditors: cm.d
 * @LastEditTime: 2021-11-06 00:58:03
 */

package myleetcode

func missingNumber(nums []int) int {
	max := 0
	min := -1
	sum1 := 0
	for _, j := range nums {
		if j >= max {
			max = j
		}
		if j == 0 {
			min = 0
		}
		sum1 = sum1 + j
	}
	for i := 0; i <= max; i++ {
		sum1 = sum1 - i
	}
	if min != 0 {
		return 0
	}
	if sum1 == 0 {
		return max + 1
	}
	return -sum1
}

func MissingNumber(nums []int) int {
	return missingNumber(nums)
}
