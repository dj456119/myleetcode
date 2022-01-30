/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2022-01-30 10:30:55
 * @LastEditors: cm.d
 * @LastEditTime: 2022-01-30 10:33:30
 */
package myleetcode

func findFinalValue(nums []int, original int) int {
	temp := make(map[int]int)
	for i, j := range nums {
		temp[j] = i
	}
	_, ok := temp[original]
	prev := original
	for ok {
		prev = original
		original = original * 2

		_, ok = temp[original]
	}
	return prev
}
