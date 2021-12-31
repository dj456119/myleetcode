/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2021-12-31 17:18:26
 * @LastEditors: cm.d
 * @LastEditTime: 2021-12-31 17:21:36
 */
package myleetcode

func countBits(n int) []int {
	result := []int{}
	for i := 0; i <= n; i++ {
		result = append(result, count(n))
	}
	return result
}

func count(n int) int {
	result := 0
	for {
		if n == 0 {
			break
		}
		n = n & (n - 1)
		if n != 0 {
			result++
		}
	}
	return result
}
