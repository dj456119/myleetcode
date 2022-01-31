/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2022-01-31 10:15:49
 * @LastEditors: cm.d
 * @LastEditTime: 2022-01-31 10:42:51
 */

package myleetcode

func rotate(matrix [][]int) {
	updown(matrix)
	leftright(matrix)
}

func updown(matrix [][]int) {
	start := 0
	end := len(matrix) - 1
	for start < end {
		for i := range matrix[start] {
			matrix[start][i], matrix[end][i] = matrix[end][i], matrix[start][i]
		}
		start++
		end--
	}
}

func leftright(matrix [][]int) {
	for y := range matrix {
		for x := y; x < len(matrix[y]); x++ {
			matrix[y][x] = matrix[x][y]
		}
	}
}
