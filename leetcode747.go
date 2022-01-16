/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2022-01-13 00:26:53
 * @LastEditors: cm.d
 * @LastEditTime: 2022-01-13 00:29:15
 */
package myleetcode

func dominantIndex(nums []int) int {
	max := 0
	init := false
	pos := 0
	for i, j := range nums {
		if j > max || init == false {
			init = true
			max = j
			pos = i
		}
	}

	for i, j := range nums {
		if i == pos {
			continue
		}
		if j*2 > max {
			return -1
		}
	}
	return pos

}
