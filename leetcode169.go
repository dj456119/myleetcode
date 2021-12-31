/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2021-12-31 16:37:44
 * @LastEditors: cm.d
 * @LastEditTime: 2021-12-31 16:46:43
 */
func majorityElement(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	pos := 0
	if len(nums)/2 == 0 {
		pos = len(nums)/2 - 1
	} else {
		pos = len(nums) / 2
	}
	return sortArray(nums)[pos]
}
