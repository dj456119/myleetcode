/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2021-12-31 12:28:26
 * @LastEditors: cm.d
 * @LastEditTime: 2021-12-31 12:35:57
 */
package myleetcode

func strStr(haystack string, needle string) int {
	if haystack == "" && needle == "" {
		return 0
	}
	if needle == "" {
		return 0
	}
	if haystack == "" {
		return -1
	}
	for i := 0; i < len(haystack); i++ {
		if haystack[i] == needle[0] {
			count := 0
			if i+len(needle) > len(haystack) {
				return -1
			}
			for j := i; j < len(needle)+i; j++ {
				if haystack[j] != needle[j-i] {
					break
				}
				count++
			}
			if count == len(needle) {
				return i
			}
		}
	}
	return -1
}
