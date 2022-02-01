/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2022-02-01 17:36:11
 * @LastEditors: cm.d
 * @LastEditTime: 2022-02-01 17:45:34
 */
package myleetcode

func dailyTemperatures(temperatures []int) []int {
	result := make([]int, len(temperatures))
	stack := []int{}
	for i := range temperatures {
		if i == 0 {
			stack = append(stack, i)
		} else {
			for len(stack) != 0 && temperatures[i] > temperatures[stack[len(stack)-1]] {
				topIndex := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				result[topIndex] = i - topIndex
			}
			stack = append(stack, i)
		}
	}
	return result
}
