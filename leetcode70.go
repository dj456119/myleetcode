/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2022-01-16 00:37:39
 * @LastEditors: cm.d
 * @LastEditTime: 2022-01-16 00:46:31
 */
package myleetcode

func climbStairs(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	fn2, fn1, now, fn := 1, 1, 2, 0

	for {
		fn = fn1 + fn2
		if now == n {
			break
		}
		fn2, fn1 = fn1, fn
		now++
	}
	return fn
}
