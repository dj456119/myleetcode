/*
 * @Descripttion:给定一个初始元素全部为 0，大小为 m*n 的矩阵 M 以及在 M 上的一系列更新操作。

操作用二维数组表示，其中的每个操作用一个含有两个正整数 a 和 b 的数组表示，含义是将所有符合 0 <= i < a 以及 0 <= j < b 的元素 M[i][j] 的值都增加 1。

在执行给定的一系列操作后，你需要返回矩阵中含有最大整数的元素个数。
 * @version:
 * @Author: cm.d
 * @Date: 2021-11-07 21:35:13
 * @LastEditors: cm.d
 * @LastEditTime: 2021-11-07 22:01:57
*/

package myleetcode

func maxCount(m int, n int, ops [][]int) int {
	minX := -1
	minY := -1
	for _, j := range ops {
		if minX == -1 || minX >= j[0] {
			minX = j[0]
		}
		if minY == -1 || minY >= j[1] {
			minY = j[1]
		}
	}
	if minX == -1 {
		minX = m
	}
	if minY == -1 {
		minY = n
	}

	return minX * minY
}

func MaxCount(m int, n int, ops [][]int) int {
	return maxCount(m, n, ops)
}
