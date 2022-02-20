/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2022-02-20 11:05:55
 * @LastEditors: cm.d
 * @LastEditTime: 2022-02-20 11:09:00
 */
package myleetcode

func countEven(num int) int {
	result := 0
	for i := 1; i <= num; i++ {
		if chai(i) {
			result++
		}
	}
	return result
}

func chai(num int) bool {
	z := 0
	for num != 0 {
		z += num % 10
		num = num / 10
	}
	if z%2 == 0 {
		return true
	}
	return false
}
