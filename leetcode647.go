/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2022-02-06 17:42:06
 * @LastEditors: cm.d
 * @LastEditTime: 2022-02-06 18:29:24
 */
package myleetcode

func countSubstrings(s string) int {
	if len(s) == 0 || len(s) == 1 {
		return len(s)
	}
	result := 0
	for i := range []byte(s) {
		result++
		if i > 0 {
			if s[i] == s[i-1] {
				z1 := i
				z2 := i - 1
				for z2 >= 0 && z1 <= len(s)-1 && s[z1] == s[z2] {
					result++
					z1++
					z2--
				}
			}

			z1 := i + 1
			z2 := i - 1
			for z2 >= 0 && z1 <= len(s)-1 && s[z1] == s[z2] {
				result++
				z1++
				z2--
			}

		}
	}
	return result
}
