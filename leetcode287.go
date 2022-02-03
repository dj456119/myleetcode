/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2022-02-03 17:37:13
 * @LastEditors: cm.d
 * @LastEditTime: 2022-02-03 17:55:09
 */
package myleetcode

func findDuplicate(nums []int) int {
	left := 1
	right := len(nums) - 1
	mid := (left + right) / 2
	for {
		midCount := 0
		count := 0
		for _, j := range nums {
			if j <= mid {
				count++
				if j == mid {
					midCount++
				}
			}
		}
		if midCount >= 2 {
			return mid
		}
		if count > mid {
			right = mid
		} else {
			left = mid + 1
		}
		mid = (left + right) / 2
	}
}
