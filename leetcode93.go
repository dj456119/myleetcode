/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2022-01-21 11:07:23
 * @LastEditors: cm.d
 * @LastEditTime: 2022-01-22 00:35:11
 */
package myleetcode

import (
	"fmt"
	"strconv"
)

func restoreIpAddresses(s string) []string {
	return ndfs([]byte(s), 0, []string{})
}

func ndfs(s []byte, now int, result []string) []string {

	if now != len(s) && len(result) == 4 {
		return nil
	}

	if now == len(s) && len(result) == 4 {
		return []string{printIP(result)}
	}

	if now == len(s) {
		return nil
	}

	results := []string{}

	length := 3
	if s[now] == '0' {
		length = 1
	}

	for i := 0; i < length; i++ {
		if now+i >= len(s) {
			break
		}
		z := s[now : now+i+1]
		m, _ := strconv.ParseInt(string(z), 10, 32)

		if m > 255 {
			break
		}
		result1 := make([]string, len(result))
		copy(result1, result)
		result1 = append(result1, string(z))
		r := ndfs(s, now+i+1, result1)
		if r != nil {
			results = append(results, r...)
		}
	}
	return results
}

func printIP(s []string) string {
	return fmt.Sprintf("%s.%s.%s.%s", s[0], s[1], s[2], s[3])
}
