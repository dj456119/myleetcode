/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2022-01-21 01:12:03
 * @LastEditors: cm.d
 * @LastEditTime: 2022-01-21 01:17:11
 */
package myleetcode

import "sort"

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	result := [][]int{}
	for _, j := range intervals {
		if len(result) == 0 {
			result = append(result, j)
		} else {
			a := result[len(result)-1]
			if a[1] > j[1] {
				continue
			}
			if a[1] < j[0] {
				result = append(result, j)
			} else {
				a[1] = j[1]
			}
		}
	}
	return result
}
