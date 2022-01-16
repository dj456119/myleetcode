/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2022-01-08 17:12:46
 * @LastEditors: cm.d
 * @LastEditTime: 2022-01-08 17:16:57
 */
package myleetcode

import "sort"

func nextPermutation(nums []int) {
	if len(nums) < 2 {
		return
	}
	i := len(nums) - 2
	j := len(nums) - 1
	for {
		if nums[i]*10+nums[j] < nums[j]*10+nums[i] {
			nums[i], nums[j] = nums[j], nums[i]
			return
		}
		i--
		j--
		if i < 0 || j < 0 {
			break
		}
	}
	sort.Slice(nums, func(x, y int) bool {
		return nums[x] < nums[y]
	})
}
