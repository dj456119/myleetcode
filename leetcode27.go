/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2021-12-31 12:25:04
 * @LastEditors: cm.d
 * @LastEditTime: 2021-12-31 12:26:05
 */
package myleetcode

func removeElement(nums []int, val int) int {
	i := 0
	count := 0
	result := 0
	for {
		if i == len(nums) {
			break
		}
		if nums[i] == val {
			count++
		} else {
			nums[i-count] = nums[i]
			result++
		}
	}
	return result
}
