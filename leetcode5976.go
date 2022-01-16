/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2022-01-09 10:30:15
 * @LastEditors: cm.d
 * @LastEditTime: 2022-01-09 10:34:59
 */
package myleetcode

func checkValid(matrix [][]int) bool {
	n := len(matrix)

	s := make([]int, n)

	for _, j := range matrix {
		for _, m := range j {
			s[m] = 1
		}
		for _, z := range s {
			if z != 1 {
				return false
			}
		}
		s = make([]int, n)
	}
	s = make([]int, n)
	for i := range matrix {
		for j := range matrix {
			s[matrix[j][i]] = 1
		}
		for _, z := range s {
			if z != 1 {
				return false
			}
		}
		s = make([]int, n)
	}
	return true
}
