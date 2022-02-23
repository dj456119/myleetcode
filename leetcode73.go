/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2022-02-24 00:54:27
 * @LastEditors: cm.d
 * @LastEditTime: 2022-02-24 00:54:28
 */
package myleetcode

func setZeroes(matrix [][]int) {
	rowCount := false
	for i := range matrix[0] {
		if matrix[0][i] == 0 {
			rowCount = true
		}
	}
	for y := range matrix {
		if y == 0 {
			continue
		}
		for x := range matrix[y] {
			if x == 0 {
				if matrix[y][x] == 0 {
					matrix[0][0] = 0
				}
			} else {
				if matrix[y][x] == 0 {
					matrix[0][x] = 0
					matrix[y][0] = 0
				}
			}
		}
	}

	for y := range matrix {
		if y == 0 {
			continue
		}
		for x := range matrix[y] {
			if x == 0 {
				continue
			}
			if matrix[y][0] == 0 || matrix[0][x] == 0 {
				matrix[y][x] = 0
			}
		}
	}

	for i := len(matrix) - 1; i >= 0; i-- {
		if matrix[0][0] == 0 {
			matrix[i][0] = 0
		}
	}

	for i := range matrix[0] {
		if rowCount {
			matrix[0][i] = 0
		}
	}

}
