/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2022-01-31 19:40:25
 * @LastEditors: cm.d
 * @LastEditTime: 2022-01-31 19:54:02
 */
package myleetcode

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	startx := len(matrix[0]) - 1
	starty := 0

	for startx >= 0 && starty < len(matrix) {
		if matrix[starty][startx] == target {
			return true
		}
		for matrix[starty][startx] < target {
			starty++
			if starty == len(matrix) {
				return false
			}
		}
		for matrix[starty][startx] > target {
			startx--
			if startx == -1 {
				return false
			}
		}
	}
	return false
}
