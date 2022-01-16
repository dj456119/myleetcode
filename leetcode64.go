/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2022-01-16 16:59:43
 * @LastEditors: cm.d
 * @LastEditTime: 2022-01-16 17:13:04
 */
package myleetcode

func minPathSum(grid [][]int) int {
	arr := make([][]int, len(grid))
	l := len(grid[0])
	for i := range arr {
		arr[i] = make([]int, l)
	}
	for y := range arr {
		for x := range arr[y] {
			if y == 0 {
				if x == 0 {
					arr[y][x] = grid[y][x]
				} else {
					arr[y][x] = arr[y][x-1] + grid[y][x]
				}
			} else {
				if x == 0 {
					arr[y][x] = arr[y-1][x] + grid[y][x]
				} else {
					arr[y][x] = min(arr[y-1][x], arr[y][x-1]) + grid[y][x]
				}
			}
		}

	}
	return arr[len(arr)-1][len(arr[0])-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
