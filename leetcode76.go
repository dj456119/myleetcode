/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2022-02-19 17:30:08
 * @LastEditors: cm.d
 * @LastEditTime: 2022-02-19 18:54:57
 */

package myleetcode

func minWindow(s string, t string) string {
	m2 := make(map[byte]int)
	m3 := make(map[byte]int)
	for i := range []byte(t) {
		m2[t[i]]++
	}
	right := 0
	left := 0
	result := 9999999999
	rs := ""
	m4 := make(map[byte]int)
	for right != len(s) {
		for right != len(s) {
			if _, ok := m2[s[right]]; ok {
				m3[s[right]]++
				if _, ok := m3[s[right]]; ok {

					if m3[s[right]] >= m2[s[right]] {
						m4[s[right]]++
					}
					if len(m4) == len(m2) {
						right++
						if right-left < result {
							result = right - left
							rs = s[left:right]
						}
						break
					}
				}
			}
			right++
		}
		if len(m4) != len(m2) {
			break
		}
		for left != right {
			if _, ok := m3[s[left]]; ok {
				if m3[s[left]] >= m2[s[left]] {
					if right-left < result {
						result = right - left
						rs = s[left:right]
					}
				}
				m3[s[left]]--
				if m3[s[left]] >= m2[s[left]] {
					if right-left < result {
						result = right - left
						rs = s[left:right]
					}
					left++
					continue
				}
				delete(m4, s[left])
				left++
				break
			}
			left++
		}
	}
	return rs
}
