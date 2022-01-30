/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2022-01-30 00:48:53
 * @LastEditors: cm.d
 * @LastEditTime: 2022-01-30 01:03:41
 */
package myleetcode

func sortColors(nums []int) {
	if len(nums) == 0 {
		return
	}
	p0 := 0
	p1 := 0

	for i, j := range nums {
		if j == 0 {
			nums[i], nums[p0] = nums[p0], nums[i]
			if p0 < p1 {
				nums[i], nums[p1] = nums[p1], nums[i]
			}
			p0++
			p1++
		}
		if j == 1 {
			nums[i], nums[p1] = nums[p1], nums[i]
			p1++
		}
	}
}
