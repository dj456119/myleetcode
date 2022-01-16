/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2022-01-15 00:18:05
 * @LastEditors: cm.d
 * @LastEditTime: 2022-01-15 00:33:34
 */
package myleetcode

func totalMoney(n int) int {
	a := n % 7
	b := n / 7
	result := 0
	for i := 1; i <= b; i++ {
		result = result + 7*i + 7*(7-1)/2
	}
	result = result + (b+1)*a + a*(a-1)/2

	return result
}
