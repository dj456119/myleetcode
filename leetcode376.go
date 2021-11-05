/*
 * @Descripttion:给定一个 正整数 num ，编写一个函数，如果 num 是一个完全平方数，则返回 true ，否则返回 false 。

进阶：不要 使用任何内置的库函数，如  sqrt 。

 * @version:
 * @Author: cm.d
 * @Date: 2021-11-04 23:05:50
 * @LastEditors: cm.d
 * @LastEditTime: 2021-11-04 23:28:03
*/

package myleetcode

func isPerfectSquare(num int) bool {
	if num == 0 || num == 1 {
		return true
	}
	start := 0
	end := num
	for {
		pos := (start + end) / 2
		if start+1 == end || start == end {
			return pos*pos == num
		}
		if pos*pos > num {
			end = pos
		} else if pos*pos < num {
			start = pos
		} else {
			return true
		}
	}
}

func IsPerfectSquare(num int) bool {
	return isPerfectSquare(num)
}
