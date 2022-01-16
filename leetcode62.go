/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2022-01-16 16:52:56
 * @LastEditors: cm.d
 * @LastEditTime: 2022-01-16 16:55:24
 */
package myleetcode

func uniquePaths(m int, n int) int {
	arr := make([][]int, m)
	for i := range arr {
		arr[i] = make([]int, n)
	}
	for y := range arr {
		if y == 0 {
			for x := range arr {
				arr[y][x] = 1
			}
		} else {
			for x := range arr {
				if x == 0 {
					arr[y][x] = 1
				} else {
					arr[y][x] = arr[y-1][x] + arr[y][x-1]
				}
			}
		}
	}
	return arr[m-1][n-1]
}
