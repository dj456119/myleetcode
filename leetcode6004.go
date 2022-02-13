/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2022-02-13 10:31:47
 * @LastEditors: cm.d
 * @LastEditTime: 2022-02-13 10:32:54
 */
package myleetcode

func countOperations(num1 int, num2 int) int {
	count := 0
	for num1 != 0 && num2 != 0 {
		if num1 < num2 {
			num2 = num2 - num1
		} else {
			num1 = num1 - num2
		}
		count++
	}
	return count
}
