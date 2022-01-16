/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2022-01-16 17:15:07
 * @LastEditors: cm.d
 * @LastEditTime: 2022-01-16 17:22:14
 */
package myleetcode

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])
	arr := make([][]int, m)
	for i := range arr {
		arr[i] = make([]int, n)
	}
	for y := range arr {
		if y == 0 {
			for x := range arr[y] {
				if obstacleGrid[y][x] == 1 {
					break
				} else {
					arr[y][x] = 1
				}
			}
		} else {
			for x := range arr[y] {
				if obstacleGrid[y][x] == 1 {
					continue
				}
				if x == 0 {
					if arr[y-1][x] == 0 {
						arr[y][x] = 0
						continue
					}
					arr[y][x] = 1
				} else {
					if arr[y-1][x] == 0 && arr[y][x-1] == 0 {
						arr[y][x] = 0
						continue
					}
					arr[y][x] = arr[y-1][x] + arr[y][x-1]
				}
			}
		}
	}
	return arr[m-1][n-1]
}
