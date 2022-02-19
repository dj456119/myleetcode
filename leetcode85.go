/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2022-02-19 00:33:26
 * @LastEditors: cm.d
 * @LastEditTime: 2022-02-19 13:11:55
 */
package myleetcode

import "fmt"

func maximalRectangle(matrix [][]byte) int {
	dparr := make([][][]int, len(matrix))
	for i := range matrix {
		dparr[i] = make([][]int, len(matrix[i]))
		for j := range dparr[i] {
			dparr[i][j] = make([]int, 2)
		}
	}
	for y := range matrix {
		for x := range matrix[y] {
			if matrix[y][x] == '0' {
				continue
			}
			if y == 0 && x == 0 {
				dparr[y][x][0] = int(matrix[y][x] - '0')
				if dparr[y][x][0] != 0 {
					dparr[y][x][1] = 1
				}
				continue
			}
			if y == 0 {
				if matrix[y][x-1] == '0' {
					dparr[y][x][0] = int(matrix[y][x] - '0')
				} else {
					dparr[y][x][0] = dparr[y][x-1][0] + int(matrix[y][x]-'0')
				}
				if matrix[y][x] != '0' {
					dparr[y][x][1] = 1
				}
				continue
			}
			if x == 0 {
				if matrix[y-1][x] == '0' {
					dparr[y][x][1] = int(matrix[y][x] - '0')
				} else {
					dparr[y][x][1] = dparr[y-1][x][1] + int(matrix[y][x]-'0')
				}
				if matrix[y][x] != '0' {
					dparr[y][x][0] = 1
				}
				continue
			}
			if dparr[y-1][x][0] == 0 {
				dparr[y][x][0] = dparr[y][x-1][0] + 1
				dparr[y][x][1] = 1
				continue
			} else if dparr[y][x-1][0] == 0 {
				dparr[y][x][0] = 1
				dparr[y][x][1] = dparr[y-1][x][1] + 1
				continue
			}
			z1 := dparr[y][x-1][0] + 1
			z2 := dparr[y-1][x][0]

			z3 := dparr[y][x-1][1]
			z4 := dparr[y-1][x][1] + 1

			m1 := min(z1, z2) * min(z3, z4)
			m2 := dparr[y][x-1][0] + 1
			m3 := dparr[y-1][x][1] + 1
			if m1 >= m2 && m1 >= m3 {
				dparr[y][x][0] = min(z1, z2)
				dparr[y][x][1] = min(z3, z4)
			}
			if m2 >= m1 && m2 >= m3 {
				dparr[y][x][0] = m2
				dparr[y][x][1] = 1
			}
			if m3 >= m1 && m3 >= m2 {
				dparr[y][x][0] = 1
				dparr[y][x][1] = m3
			}
		}
		fmt.Println(dparr[y])
	}
	max := -1
	for y := range dparr {
		for x := range dparr[y] {
			z := dparr[y][x][0] * dparr[y][x][1]
			if z > max {
				max = z
			}
		}
	}
	return max
}

func maximalRectangle1(matrix [][]byte) int {
	dparr := make([][][]int, len(matrix))
	for i := range matrix {
		dparr[i] = make([][]int, len(matrix[i]))
		for j := range dparr[i] {
			dparr[i][j] = make([]int, 2)
		}
	}
	for y := range matrix {
		for x := range matrix[y] {
			if matrix[y][x] == '0' {
				continue
			}
			if y == 0 && x == 0 {
				dparr[y][x][0] = int(matrix[y][x] - '0')
				if dparr[y][x][0] != 0 {
					dparr[y][x][1] = 1
				}
				continue
			}
			if y == 0 {
				if matrix[y][x-1] == '0' {
					dparr[y][x][0] = int(matrix[y][x] - '0')
				} else {
					dparr[y][x][0] = dparr[y][x-1][0] + int(matrix[y][x]-'0')
				}
				if matrix[y][x] != '0' {
					dparr[y][x][1] = 1
				}
				continue
			}
			if x == 0 {
				if matrix[y-1][x] == '0' {
					dparr[y][x][1] = int(matrix[y][x] - '0')
				} else {
					dparr[y][x][1] = dparr[y-1][x][1] + int(matrix[y][x]-'0')
				}
				if matrix[y][x] != '0' {
					dparr[y][x][0] = 1
				}
				continue
			}
			dparr[y][x][0] = dparr[y][x-1][0] + 1
			dparr[y][x][1] = dparr[y-1][x][1] + 1
		}
		fmt.Println(dparr[y])
	}
	max := -1
	for y := range dparr {
		for x := range dparr[y] {
			height := dparr[y][x][1]
			for m := y - height; m < height; m++ {
				z := dparr[x][m][0] * height
				if z > max {
					max = z
				}
				height--
			}

		}
	}
	return max
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
