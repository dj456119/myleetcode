/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2021-12-31 12:36:36
 * @LastEditors: cm.d
 * @LastEditTime: 2021-12-31 12:47:05
 */
package myleetcode

func lengthOfLastWord(s string) int {
	lastpos := 0
	for i, j := range s {
		if i == len(s)-1 {
			if j != ' ' {
				lastpos = i
			}
			break
		}
		if s[i+1] == ' ' && j != ' ' {
			lastpos = i
		}
	}
	result := 0
	for lastpos >= 0 {
		if s[lastpos] != ' ' {
			result++
		} else {
			return result
		}
		lastpos--
	}
	return result

}
