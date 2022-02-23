/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2022-02-24 00:56:12
 * @LastEditors: cm.d
 * @LastEditTime: 2022-02-24 01:21:35
 */
package myleetcode

func searchMatrix(matrix [][]int, target int) bool {
	start := 0
	end := len(matrix)*len(matrix[0]) - 1
	mid := (end + start) / 2
	midY := mid / len(matrix[0])
	midX := mid % len(matrix[0])
	for {
		if matrix[midY][midX] == target {
			return true
		}
		if start >= end {
			return false
		}
		if matrix[midY][midX] < target {
			start = mid + 1
		} else {
			end = mid
		}
		mid = (end + start) / 2
		midY = mid / len(matrix[0])
		midX = mid % len(matrix[0])
	}
}
