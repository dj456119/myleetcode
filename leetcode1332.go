/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2022-01-22 00:14:27
 * @LastEditors: cm.d
 * @LastEditTime: 2022-01-22 00:14:45
 */
package myleetcode

func removePalindromeSub(s string) int {
	for i, n := 0, len(s); i < n/2; i++ {
		if s[i] != s[n-1-i] {
			return 2
		}
	}
	return 1
}
