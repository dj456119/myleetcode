/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2022-01-22 16:50:04
 * @LastEditors: cm.d
 * @LastEditTime: 2022-01-22 16:52:43
 */
package myleetcode

func isPalindrome(s string) bool {
	buff := []byte{}
	for _, j := range []byte(s) {
		if j >= '0' && j <= '9' || j >= 'a' && j <= 'z' {
			buff = append(buff, j)
		} else if j >= 'A' && j <= 'Z' {
			buff = append(buff, j+('a'-'A'))
		}
	}
	if len(buff) == 0 {
		return true
	}
	i := 0
	j := len(buff) - 1
	for {
		if i >= j {
			return true
		}
		if buff[i] != buff[j] {
			return false
		}
		i++
		j--
	}

}
