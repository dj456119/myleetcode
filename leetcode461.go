/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2021-12-31 17:43:34
 * @LastEditors: cm.d
 * @LastEditTime: 2021-12-31 23:46:03
 */
package myleetcode

func hammingDistance(x int, y int) int {
	result := 0
	for {
		if x == 0 && y == 0 {
			return result
		}
		x1 := x % 2
		y1 := y % 2
		if x1 = y1 {
			result++
		}
		x = x / 2
		y = y / 2
	}
}
