/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2022-02-04 23:56:55
 * @LastEditors: cm.d
 * @LastEditTime: 2022-02-05 01:43:03
 */
package myleetcode

func longestPalindrome(s string) string {
	max := 0
	result := string(s[0])
	for i := range []byte(s) {
		length := 0
		m := i
		n := i
		if m > 0 {
			m--
			for {
				if s[m] != s[n] {
					break
				}
				length = length + 2
				if length > max {
					max = length
					result = s[m : n+1]
				}
				if m-1 >= 0 {
					m--
				} else {
					break
				}
				if n+1 < len(s) {
					n++
				} else {
					break
				}
			}

		}
		length = -1
		m = i
		n = i
		for {
			if s[m] != s[n] {
				break
			}
			length = length + 2
			if length > max {
				max = length
				result = s[m : n+1]
			}
			if m-1 >= 0 {
				m--
			} else {
				break
			}
			if n+1 < len(s) {
				n++
			} else {
				break
			}
		}
	}
	return result
}
