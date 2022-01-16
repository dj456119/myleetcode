/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2022-01-07 00:28:04
 * @LastEditors: cm.d
 * @LastEditTime: 2022-01-07 00:30:12
 */
package myleetcode

func maxDepth(s string) int {
	stack := []int{}
	result := 0
	now := 0
	for _, j := range []byte(s) {
		switch j {
		case '(':
			stack = append(stack, '(')
			now++
			if now > result {
				result = now
			}
		case ')':
			stack = stack[:len(stack)-1]
			now--
		default:
		}
	}
	return result
}
