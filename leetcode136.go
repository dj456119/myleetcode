/*
 * @Descripttion:给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。
 * @version:
 * @Author: cm.d
 * @Date: 2021-11-06 11:59:43
 * @LastEditors: cm.d
 * @LastEditTime: 2021-11-06 12:00:48
 */

package myleetcode

func singleNumber136(nums []int) int {
	result := 0
	for _, j := range nums {
		result = result ^ j
	}
	return result
}

func SingleNumber136(nums []int) int {
	return singleNumber136(nums)
}
