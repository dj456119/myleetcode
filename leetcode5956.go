/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2021-12-19 12:44:25
 * @LastEditors: cm.d
 * @LastEditTime: 2021-12-19 12:44:25
 */
package myleetcode

func firstPalindrome(words []string) string {
	for _, word := range words {
		i := 0
		j := len(word) - 1
		for {
			if j <= i {
				return word
			}
			if word[i] != word[j] {
				break
			}
			i++
			j--
		}
	}
	return ""
}
