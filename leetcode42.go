/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2022-02-13 20:26:06
 * @LastEditors: cm.d
 * @LastEditTime: 2022-02-13 22:36:43
 */
package myleetcode

func trap(height []int) int {
	dparrLeft := make([]int, len(height))
	dparrRight := make([]int, len(height))
	max := -1
	for i := range height {
		if height[i] > max {
			max = height[i]
		}
		dparrLeft[i] = max
	}
	max = -1
	for i := len(height) - 1; i >= 0; i-- {
		if height[i] > max {
			max = height[i]
		}
		dparrRight[i] = max
	}
	result := 0
	for i := range height {
		result = result + (min42(dparrLeft[i], dparrRight[i]) - height[i])
	}
	return result
}

func min42(a, b int) int {
	if a < b {
		return a
	}
	return b
}
