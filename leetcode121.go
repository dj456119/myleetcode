/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2022-01-16 16:29:29
 * @LastEditors: cm.d
 * @LastEditTime: 2022-01-16 16:33:06
 */
package myleetcode

func maxProfit(prices []int) int {
	m := 0
	result := 0
	for i, j := range prices {
		if i == 0 {
			m = j
			result = 0
		} else {
			m = min(j, m)
			result = max(result, j-m)
		}
	}
	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
