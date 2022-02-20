/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2022-02-19 22:38:00
 * @LastEditors: cm.d
 * @LastEditTime: 2022-02-19 22:42:21
 */
package myleetcode

func sumOfThree(num int64) []int64 {
	if num%3 != 2 {
		return []int64{}
	}
	z := num / 3
	result := []int64{}
	result = append(result, z-1)
	result = append(result, z)
	result = append(result, z+1)

	return result
}
