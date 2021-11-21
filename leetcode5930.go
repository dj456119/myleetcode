/*
 * @Descripttion:街上有 n 栋房子整齐地排成一列，每栋房子都粉刷上了漂亮的颜色。给你一个下标从 0 开始且长度为 n 的整数数组 colors ，其中 colors[i] 表示第  i 栋房子的颜色。

返回 两栋 颜色 不同 房子之间的 最大 距离。

第 i 栋房子和第 j 栋房子之间的距离是 abs(i - j) ，其中 abs(x) 是 x 的绝对值。

 * @version:
 * @Author: cm.d
 * @Date: 2021-11-22 00:44:24
 * @LastEditors: cm.d
 * @LastEditTime: 2021-11-22 00:44:24
*/
package myleetcode

func maxDistance(colors []int) int {
	max := 0
	for i, j := range colors {
		for m, n := range colors {
			if i == m {
				continue
			}
			if j != n {
				if m-i > max {
					max = m - i
				}
			}
		}
	}
	return max
}
