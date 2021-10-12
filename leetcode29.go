/*
 * @Descripttion:leetcode29，给定两个整数，被除数 dividend 和除数 divisor。将两数相除，要求不使用乘法、除法和 mod 运算符。返回被除数 dividend 除以除数 divisor 得到的商。
 * @version:
 * @Author: cm.d
 * @Date: 2021-10-12 17:56:30
 * @LastEditors: cm.d
 * @LastEditTime: 2021-10-12 18:51:30
 */

package myleetcode

func divide(dividend int, divisor int) int {
	INT_MAX := int(int32(^uint32(0) >> 1))
	INT_MIN := ^INT_MAX
	if dividend == INT_MIN {
		if divisor == 1 {
			return INT_MIN
		}
		if divisor == -1 {
			return INT_MAX
		}
	}
	if divisor == INT_MIN {
		if dividend == INT_MIN {
			return 1
		}
		return 0
	}

	if dividend == 0 {
		return 0
	}
	result := 0
	//false为正，true为负
	f := false
	if dividend < 0 {
		dividend = -dividend
		f = !f
	}
	if divisor < 0 {
		divisor = -divisor
		f = !f
	}
	if dividend < divisor {
		return 0
	}
	for dividend >= divisor {
		dividend = dividend - divisor
		result++
	}
	if !f {
		return result
	}
	return -result
}

func Divide(dividend int, divisor int) int {
	return divide(dividend, divisor)
}
