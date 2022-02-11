/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2022-02-12 00:05:35
 * @LastEditors: cm.d
 * @LastEditTime: 2022-02-12 00:37:33
 */
package myleetcode

func generateMatrix(n int) [][]int {
	x1 := 0
	x2 := n - 1

	y1 := 0
	y2 := n - 1

	result := make([][]int, n)
	t := 1
	for i := range result {
		result[i] = make([]int, n)
	}
	for {
		if x1 == x2 {
			result[y1][x1] = t
			break
		}
		if x1 > x2 {
			break
		}
		for i := x1; i < x2; i++ {
			result[y1][i] = t
			t++
		}
		for i := y1; i < y2; i++ {
			result[i][x2] = t
			t++
		}
		for i := x2; i > x1; i-- {
			result[y2][i] = t
			t++
		}
		for i := y2; i < y1; i-- {
			result[i][x1] = t
			t++
		}
		x1++
		y1++
		x2--
		y2--
	}
	return result
}
