/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2022-02-06 10:31:45
 * @LastEditors: cm.d
 * @LastEditTime: 2022-02-06 10:36:55
 */
package myleetcode

import "sort"

func sortEvenOdd(nums []int) []int {
	ji := []int{}
	ou := []int{}
	for i, j := range nums {
		if i%2 == 0 {
			ou = append(ou, j)
		} else {
			ji = append(ji, j)
		}
	}
	sort.Slice(ji, func(i, j int) bool {
		return ji[i] > ji[j]
	})
	sort.Slice(ou, func(i, j int) bool {
		return ou[i] < ou[j]
	})
	result := []int{}
	for z := range ou {
		result = append(result, ou[z])
		if z < len(ji) {
			result = append(result, ji[z])
		}
		z++
	}
	return result
}
