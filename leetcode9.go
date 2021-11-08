/*
 * @Descripttion:给你一个整数 x ，如果 x 是一个回文整数，返回 true ；否则，返回 false 。

回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。例如，121 是回文，而 123 不是。

 * @version:
 * @Author: cm.d
 * @Date: 2021-11-09 00:30:03
 * @LastEditors: cm.d
 * @LastEditTime: 2021-11-09 00:41:49
*/
package myleetcode

import "fmt"

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	sx := fmt.Sprintf("%d", x)
	i := 0
	j := len(sx) - 1
	for {
		if i == j {
			return true
		}
		if i+1 == j {
			return sx[i] == sx[j]
		}
		if sx[i] != sx[j] {
			return false
		}
		i++
		j--
	}

}

func IsPalindrome(x int) bool {
	return isPalindrome(x)
}
