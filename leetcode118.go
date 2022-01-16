/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2022-01-16 17:54:45
 * @LastEditors: cm.d
 * @LastEditTime: 2022-01-16 18:07:00
 */
package myleetcode

func generate(numRows int) [][]int {
	result := [][]int{}
	if numRows == 0 {
		return result
	}
	result = append(result, []int{1})
	if numRows == 1 {
		return result
	}
	result = append(result, []int{1, 1})
	if numRows == 2 {
		return result
	}
	for i := 3; i <= numRows; i++ {
		s := make([]int, i)
		s[0] = 1
		s[len(s)-1] = 1
		for j := 1; j <= len(s)-2; j++ {
			s[j] = result[i-2][j-1] + result[i-2][j]
		}
		result = append(result, s)
	}
	return result
}
