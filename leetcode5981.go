/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2022-01-30 10:36:50
 * @LastEditors: cm.d
 * @LastEditTime: 2022-01-30 10:47:27
 */
package myleetcode

import "fmt"

func maxScoreIndices(nums []int) []int {
	count1 := 0
	for _, j := range nums {
		if j == 1 {
			count1++
		}
	}
	if count1 == 0 {
		return []int{len(nums)}
	}
	if count1 == len(nums) {
		return []int{0}
	}
	left := 0
	right := count1
	max := count1
	for _, j := range nums {
		if j == 0 {
			left++
		} else {
			right--
		}
		if left+right > max {
			max = left + right
		}
	}
	left = 0
	right = count1
	result := []int{}
	if count1 == max {
		result = append(result, 0)
	}
	for i, j := range nums {
		if j == 0 {
			left++
		} else {
			right--
		}
		if left+right == max {
			fmt.Println(i, max)
			result = append(result, i+1)
		}
	}
	return result
}
