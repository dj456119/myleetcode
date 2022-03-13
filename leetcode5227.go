/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2022-03-13 11:17:36
 * @LastEditors: cm.d
 * @LastEditTime: 2022-03-13 13:03:22
 */
package myleetcode

func maximumTop(nums []int, k int) int {
	if len(nums) == 1 && k%2 != 0 {
		return -1
	}
	if len(nums) == 1 && k%2 == 0 {
		return nums[0]
	}
	max := -9999999999
	for i := range nums {
		if i+1 >= k {
			break
		}
		if nums[i] > max {
			max = nums[i]
		}
	}
	if k <= len(nums)-1 && max < nums[k] {
		return nums[k]
	}
	return max
}
