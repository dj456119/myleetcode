/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2022-02-25 01:10:55
 * @LastEditors: cm.d
 * @LastEditTime: 2022-02-25 01:10:56
 */
package myleetcode

func removeDuplicates(nums []int) int {
	step := 0
	nowCount := 1
	nowInt := nums[0]
	for i := range nums {
		if i == 0 {
			continue
		}
		if nums[i] == nowInt && nowCount == 2 {
			step++
		} else if nums[i] == nowInt && nowCount != 2 {
			nowCount++
			nums[i], nums[i-step] = nums[i-step], nums[i]
		} else {
			nowCount = 1
			nowInt = nums[i]
			nums[i], nums[i-step] = nums[i-step], nums[i]
		}
	}
	return len(nums) - step
}
