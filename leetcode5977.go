/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2022-01-09 11:12:44
 * @LastEditors: cm.d
 * @LastEditTime: 2022-01-10 22:30:13
 */
package myleetcode

func minSwaps(nums []int) int {
	maxLengthFirst := 0

	count := 0
	for _, j := range nums {
		if j == 1 {
			count++
		}
	}
	afterzcount := 0
	if nums[0] == 1 {
		last := 0
		for z := len(nums) - 1; z >= 0; z-- {
			if nums[z] != 1 {
				first = last
				break
			}
		}
	} else {
		for i, j := range nums {
			if j == 1 {
				first = i
				break
			}
		}
	}
	return result
}
