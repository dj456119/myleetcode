/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2022-02-20 15:46:10
 * @LastEditors: cm.d
 * @LastEditTime: 2022-02-20 15:53:20
 */
package myleetcode

func minDistance(word1 string, word2 string) int {
	dparr := make([][]int, len(word1)+1)
	for i := range dparr {
		dparr[i] = make([]int, len(word2)+1)
	}
	result := 99999999999
	for y := range dparr {
		for x := range dparr[y] {
			if y == 0 {
				dparr[y][x] = x
				continue
			}
			if x == 0 {
				dparr[y][x] = y
				continue
			}
			dparr[y][x] = min(dparr[y-1][x], min(dparr[y-1][x-1], dparr[y][x-1])) + 1
			if x == len(dparr[y])-1 {
				if dparr[y][x] < result {
					result = dparr[y][x]
				}
			}
		}
	}
	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
