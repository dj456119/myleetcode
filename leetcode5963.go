/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2021-12-26 10:50:03
 * @LastEditors: cm.d
 * @LastEditTime: 2021-12-31 10:36:51
 */
package myleetcode

import (
	"fmt"
	"strconv"
)

func isSameAfterReversals(num int) bool {
	n1 := fmt.Sprintf("%d", num)
	n2 := reverseString(n1)
	n3, _ := strconv.ParseInt(n2, 10, 32)
	n4 := fmt.Sprintf("%d", n3)
	n5 := reverseString(n4)
	return n1 == n5
}

func reverseString(s string) string {
	runes := []rune(s)
	for from, to := 0, len(runes)-1; from < to; from, to = from+1, to-1 {
		runes[from], runes[to] = runes[to], runes[from]
	}
	return string(runes)
}
