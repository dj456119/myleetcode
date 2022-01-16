/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2022-01-08 16:49:14
 * @LastEditors: cm.d
 * @LastEditTime: 2022-01-08 16:56:16
 */
package myleetcode

func generateParenthesis(n int) []string {
	return digui(0, n, 0, []byte{})
}

func digui(now int, n int, nowleft int, result []byte) []string {
	if now == n*2 {
		if nowleft == 0 {
			return []string{string(result)}
		}
		return nil
	}
	result1 := []string{}
	if nowleft < n {
		r1 := make([]byte, len(result))
		copy(r1, result)
		r1 = append(r1, '(')
		r2 := digui(now+1, n, nowleft+1, r1)
		if r2 != nil {
			result1 = append(result1, r2...)
		}
	}
	if nowleft > 0 {
		r1 := make([]byte, len(result))
		copy(r1, result)
		r1 = append(r1, ')')
		r2 := digui(now+1, n, nowleft-1, r1)
		if r2 != nil {
			result1 = append(result1, r2...)
		}
	}
	return result1
}
