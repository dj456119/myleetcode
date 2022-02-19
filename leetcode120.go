/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2022-02-18 21:31:42
 * @LastEditors: cm.d
 * @LastEditTime: 2022-02-18 21:31:42
 */
package myleetcode

func minimumTotal(triangle [][]int) int {
	if len(triangle) == 0 {
		return -1
	}
	dparr := make([][]int, len(triangle))
	for i := range dparr {
		dparr[i] = make([]int, len(triangle[i]))
	}
	for y := range triangle {
		for x := range triangle[y] {
			if x == 0 && y == 0 {
				dparr[0][0] = triangle[y][x]
				continue
			}
			if x == 0 {
				dparr[y][x] = dparr[y-1][x] + triangle[y][x]
				continue
			}
			if x == len(triangle[y])-1 {
				dparr[y][x] = dparr[y-1][len(triangle[y-1])-1] + triangle[y][x]
				continue
			}
			dparr[y][x] = min(dparr[y-1][x-1], dparr[y-1][x]) + triangle[y][x]
		}
		//fmt.Println(dparr[y])
	}

	min := 9999999999
	for i := range dparr[len(dparr)-1] {
		if dparr[len(dparr)-1][i] < min {
			min = dparr[len(dparr)-1][i]
		}
	}
	return min
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
