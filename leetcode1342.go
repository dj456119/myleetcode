/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2022-01-31 01:05:55
 * @LastEditors: cm.d
 * @LastEditTime: 2022-01-31 01:07:00
 */
package myleetcode

func numberOfSteps(num int) int {
	if num == 0 {
		return 0
	}
	result := 0
	for num != 0 {
		if num%2 == 0 {
			result++
			num = num / 2
		} else {
			result++
			num--
		}

	}
	return result
}
