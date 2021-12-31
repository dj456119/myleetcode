/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2021-12-31 17:30:02
 * @LastEditors: cm.d
 * @LastEditTime: 2021-12-31 17:33:32
 */
package myleetcode

func findDisappearedNumbers(nums []int) []int {
	n := len(nums)
	arr := make([]bool, n)
	result := []int{}
	for _, j := range nums {
		arr[j] = true
	}
	for i, j := range arr {
		if !j {
			result = append(result, i)
		}
	}
	return result
}
