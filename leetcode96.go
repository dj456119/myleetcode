/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2022-02-06 21:25:53
 * @LastEditors: cm.d
 * @LastEditTime: 2022-02-06 21:34:32
 */
package myleetcode

func numTrees(n int) int {
	dparr := make([]int, n+1)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	dparr[0] = 1
	dparr[1] = 1
	m := n
	for n = 2; n < m+1; n++ {
		for i := 1; i <= n; i++ {
			dparr[n] = dparr[n] + dparr[i-1]*dparr[n-i]
		}
	}
	return dparr[len(dparr)-1]
}
