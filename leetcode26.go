/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2021-12-31 12:12:37
 * @LastEditors: cm.d
 * @LastEditTime: 2021-12-31 12:24:20
 */
package myleetcode

func removeDuplicates(nums []int) int {
	//当前遍历到的元素需要向前移动的位数
	count := 0
	i := 0
	result := 0
	if len(nums) == 1 {
		return 1
	}
	for {
		if i >= len(nums)-1 {
			break
		}
		if nums[i] == nums[i+1] {
			count++
		} else {
			nums[i-count] = nums[i]
			result++
		}
		i++
	}
	if len(nums) > 1 {
		result++
		nums[len(nums)-1-count] = nums[len(nums)-1]
	}
	return result
}
