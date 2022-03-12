/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2022-03-06 11:46:40
 * @LastEditors: cm.d
 * @LastEditTime: 2022-03-12 13:49:24
 */
package myleetcode

import "fmt"

func cellsInRange(s string) []string {
	startC := s[0]
	s1 := int(s[1] - '0')
	startCC := s[3]
	s2 := int(s[4] - '0')
	result := []string{}
	for i := startC; i <= startCC; i++ {
		for j := s1; j <= s2; j++ {
			result = append(result, fmt.Sprintf("%c%d", i, j))
		}
	}
	return result
}
