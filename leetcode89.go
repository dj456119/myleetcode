/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2022-02-12 02:02:00
 * @LastEditors: cm.d
 * @LastEditTime: 2022-02-12 02:14:22
 */
package myleetcode

func grayCode(n int) []int {
	result := []int{0}
	m := 1
	for i := 1; i <= n; i++ {
		f := len(result)
		for z := f - 1; z >= 0; z-- {
			result = append(result, result[z]+m)
		}
		m = m * 2
	}
	return result
}
