/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2022-01-08 15:02:03
 * @LastEditors: cm.d
 * @LastEditTime: 2022-01-08 15:13:58
 */
package leetcode

func letterCombinations(digits string) []string {
	shuzu := [][]byte{{}, {}, {'a', 'b', 'c'}, {'d', 'e', 'f'}, {'g', 'h', 'i'}, {'j', 'k', 'l'}, {'m', 'n', 'o'}, {'p', 'q', 'r', 's'}, {'t', 'u', 'v'}, {'w', 'x', 'y', 'z'}}
	return digui([]byte(digits), shuzu, 0, []byte{})
}

func digui(digits []byte, shuzu [][]byte, now int, result []byte) []string {
	if now == len(digits) {
		return []string{string(result)}
	}
	x := shuzu[digits[now]-'0']
	r := []string{}
	for _, j := range x {
		r1 := make([]byte, len(result))
		copy(r1, result)
		r1 = append(r1, j)
		r2 := digui(digits, shuzu, now+1, r1)
		r = append(r, r2...)
	}
	return r
}
