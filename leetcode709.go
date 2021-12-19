/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2021-12-19 12:47:43
 * @LastEditors: cm.d
 * @LastEditTime: 2021-12-19 12:56:03
 */
package myleetcode

func toLowerCase(s string) string {
	result := make([]byte, len(s))
	for i := range s {
		if s[i] >= 'A' && s[i] <= 'Z' {
			result[i] = s[i] + ('a' - 'A')
		} else {
			result[i] = s[i]
		}
	}
	return string(result)
}
