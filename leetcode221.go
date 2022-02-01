/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2022-02-01 12:54:34
 * @LastEditors: cm.d
 * @LastEditTime: 2022-02-01 13:23:23
 */
package myleetcode

func maximalSquare(matrix [][]byte) int {
	dparr := make([][]int, len(matrix))
	for i := range matrix {
		dparr = append(dparr, make([]int, len(matrix[i])))
	}
	max := 0
	for y := range matrix {
		for x := range matrix[y] {
			if y == 0 || x == 0 {
				dparr[y][x] = int(matrix[y][x])
			} else {
				z1 := dparr[y-1][x-1]
				xmin := 0
				ymin := 0
				for i := y; i >= 0 && matrix[i][x] != 0; i-- {
					xmin++
				}
				for i := x; x >= 0 && matrix[y][i] != 0; i-- {
					ymin++
				}
				xymin := min(xmin, ymin)
				dparr[y][x] = min(xymin, z1+1)
			}
			if dparr[y][x] > max {
				max = dparr[y][x]
			}
		}
	}
	return max * max
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
