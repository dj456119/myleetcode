/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2022-01-24 00:12:08
 * @LastEditors: cm.d
 * @LastEditTime: 2022-01-24 00:43:58
 */
package myleetcode

func spiralOrder(matrix [][]int) []int {
	zbx1 := 0
	zby1 := 0
	zbx2 := len(matrix[0]) - 1
	zby2 := 0
	zbx3 := zbx2
	zby3 := len(matrix) - 1
	zbx4 := 0
	zby4 := zby3
	result := []int{}
	for {

		if zbx1 == zbx2 && zby1 == zby4 {
			result = append(result, matrix[zby1][zbx1])
			break
		}
		if zbx1 == zbx2 && zby1 != zby4 {
			for i := zby1; i <= zby4; i++ {
				result = append(result, matrix[i][zbx1])
			}
			break
		}
		if zbx1 != zbx2 && zby1 == zby4 {
			for i := zbx1; i <= zbx2; i++ {
				result = append(result, matrix[zby1][i])
			}
			break
		}

		if zbx1 > zbx2 || zby1 > zby4 {
			break
		}
		for i := zbx1; i < zbx2; i++ {
			result = append(result, matrix[zby1][i])
		}
		for i := zby2; i < zby3; i++ {
			result = append(result, matrix[i][zbx3])
		}
		for i := zbx3; i > zbx4; i-- {
			result = append(result, matrix[zby4][i])
		}
		for i := zby4; i > zby1; i-- {
			result = append(result, matrix[i][zbx1])
		}
		zbx1++
		zby1++
		zbx2--
		zby2++
		zbx3--
		zby3--
		zbx4++
		zby4--
	}
	return result
}
