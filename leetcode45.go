/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2022-01-01 15:29:33
 * @LastEditors: cm.d
 * @LastEditTime: 2022-01-02 00:49:27
 */
package myleetcode

import "fmt"

func jump(nums []int) int {
	if len(nums) == 0 || len(nums) == 1 {
		return 0
	}
	count := 0
	for i := 0; i < len(nums); {
		j := nums[i]
		if j+i >= len(nums)-1 {
			count++
			break
		}
		max := j
		max_m := i
		end := i + j
		for m := i + 1; m <= end; m++ {
			fmt.Println(m, i, j, m+j, nums[m])
			if nums[m]+m >= len(nums)-1 {
				max = nums[m] + m
				max_m = m
				break
			}
			if nums[m]+m > max {
				max_m = m
				max = nums[m] + m
			}
		}
		i = max_m
		count++
	}
	return count
}
