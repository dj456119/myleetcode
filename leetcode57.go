/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2022-01-22 17:34:10
 * @LastEditors: cm.d
 * @LastEditTime: 2022-01-22 19:24:59
 */
package myleetcode

func insert(intervals [][]int, newInterval []int) [][]int {
	left, right := newInterval[0], newInterval[1]
	merge := false
	result := [][]int{}
	for _, j := range intervals {
		if j[1] < left {
			result = append(result, j)
		} else if j[0] > right {
			if !merge {
				result = append(result, []int{left, right})
				merge = true
			}
			result = append(result, j)
		} else {
			left = min(j[0], left)
			right = max(j[1], right)
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
