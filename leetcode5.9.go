/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2022-01-16 00:47:26
 * @LastEditors: cm.d
 * @LastEditTime: 2022-01-16 00:48:35
 */
package myleetcode

func fib(n int) int {
	if n == 0 || n == 1 {
		return 1
	}

	fn2, fn1, fn, now := 1, 1, 2, 2

	for now != n {
		fn = fn1 + fn2
		fn1, fn2 = fn, fn1
	}

	return fn
}
