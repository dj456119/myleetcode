/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2022-01-15 17:23:12
 * @LastEditors: cm.d
 * @LastEditTime: 2022-01-15 17:55:21
 */
package myleetcode

func findMaxForm(strs []string, m int, n int) int {
	arr := make([][][]int, len(strs))
	for i := range arr {
		arr[i] = make([][]int, m+1)
		for j := range arr[i] {
			arr[i][j] = make([]int, n+1)
		}
	}
	strb := [][]int{}
	for _, j := range strs {
		strb = append(strb, count(j))
	}
	result := 0
	for y := range arr {
		if y == 0 {
			if strb[0][0] <= m && strb[0][1] <= n {
				arr[y][strb[0][0]][strb[0][1]]++
				result = 1
			}
			continue
		} else {
			if strb[y][0] <= m && strb[y][1] <= n {
				arr[y][strb[y][0]][strb[y][1]] = 1
			}

		}
		for x := range arr[y] {
			for z := range arr[y][x] {
				if arr[y-1][x][z] != 0 {
					if arr[y-1][x][z] > arr[y][x][z] {
						arr[y][x][z] = arr[y-1][x][z]
					}
					if x+strb[y][0] <= m && z+strb[y][1] <= n {
						arr[y][x+strb[y][0]][z+strb[y][1]] = arr[y-1][x][z] + 1
					}
				}
				if arr[y][x][z] > result {
					result = arr[y][x][z]
				}
			}
		}
	}
	return result
}

func count(str string) []int {
	a1 := 0
	a2 := 0
	for _, j := range str {
		if j == '0' {
			a1++
		} else {
			a2++
		}
	}
	return []int{a1, a2}
}
