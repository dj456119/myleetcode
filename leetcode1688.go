/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2022-01-25 23:56:56
 * @LastEditors: cm.d
 * @LastEditTime: 2022-01-25 23:58:06
 */
package myleetcode

func numberOfMatches(n int) int {
	count := 0
	for n > 1 {
		z := n / 2
		z1 := n % 2
		if z1 == 0 {
			count += n / 2
		} else {
			count += n/2 + 1
		}
		n = z
	}
	return count
}
