/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2022-01-16 18:08:15
 * @LastEditors: cm.d
 * @LastEditTime: 2022-01-16 18:15:22
 */
package myleetcode

func getRow(rowIndex int) []int {
	numRows := rowIndex
	result := [][]int{}
	result = append(result, []int{1})
	result = append(result, []int{1, 1})
	for i := 2; i <= numRows; i++ {
		s := make([]int, i+1)
		s[0] = 1
		s[len(s)-1] = 1
		for j := 1; j <= len(s)-2; j++ {
			s[j] = result[i-1][j-1] + result[i-1][j]
		}
		result = append(result, s)
	}
	return result[rowIndex]
}
