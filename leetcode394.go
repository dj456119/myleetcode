/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2022-02-03 17:57:33
 * @LastEditors: cm.d
 * @LastEditTime: 2022-02-03 18:28:25
 */
package myleetcode

import (
	"fmt"
	"strconv"
)

func decodeString(s string) string {
	stack := [][]byte{}
	leftCount := 0
	for _, j := range []byte(s) {
		if len(stack) == 0 {
			stack = append(stack, []byte{j})
			continue
		}
		if j == '[' {
			stack = append(stack, []byte{j})
			leftCount++
			continue
		}
		if j == ']' {
			zfc := stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			num := stack[len(stack)-1]
			numInt, _ := strconv.ParseInt(string(num), 10, 32)
			r := ""
			for i := 0; i < int(numInt); i++ {
				r = fmt.Sprintf("%s%s", r, string(zfc))
			}
			stack[len(stack)-1] = []byte(r)
			if len(stack) > 1 {
				if stack[len(stack)-2][0] != '[' {
					zz := stack[len(stack)-2]
					zz = append(zz, []byte(r)...)
					stack = stack[:len(stack)-1]
					stack[len(stack)-1] = zz
				}
			}
			continue
		}
		if isNum(j) {
			if isNum(stack[len(stack)-1][0]) {
				z := stack[len(stack)-1]
				z = append(z, j)
				stack[len(stack)-1] = z
			} else {
				stack = append(stack, []byte{j})
			}
			continue
		}
		if isZM(j) {
			if isZM(stack[len(stack)-1][0]) {
				z := stack[len(stack)-1]
				z = append(z, j)
				stack[len(stack)-1] = z
			} else {
				stack = append(stack, []byte{j})
			}
		}
	}
	result := ""
	for _, j := range stack {
		result = fmt.Sprintf("%s%s", result, string(j))
	}
	return result
}

func isNum(b byte) bool {
	if b >= '0' && b <= '9' {
		return true
	}
	return false
}

func isZM(b byte) bool {
	if b >= 'a' && b <= 'z' {
		return true
	}
	return false
}
