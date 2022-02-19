/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2022-02-19 15:06:01
 * @LastEditors: cm.d
 * @LastEditTime: 2022-02-19 15:10:20
 */
package myleetcode

import "fmt"

func longestValidParentheses(s string) int {
	stc := []int{}
	que := []int{}
	for i := range []byte(s) {
		if len(s) == 0 {
			stc = append(stc, i)
			continue
		}
		if s[i] == ')' && s[stc[len(stc)-1]] == '(' {
			que = append(que, stc[len(stc)-1])
			que = append(que, i)
			stc = stc[:len(stc)-1]
		} else {
			stc = append(stc, i)
		}
	}
	fmt.Println(que, stc)
	return -1
}
