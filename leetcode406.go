/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2022-02-01 10:43:14
 * @LastEditors: cm.d
 * @LastEditTime: 2022-02-01 10:48:42
 */
package myleetcode

import "sort"

func reconstructQueue(people [][]int) [][]int {
	sort.Slice(people, func(i, j int) bool {
		if people[i][0] < people[j][0] {
			return false
		} else if people[i][0] == people[j][0] {
			if people[i][1] < people[j][1] {
				return false
			} else {
				return true
			}
		} else {
			return true
		}
	})
	for i, j := range people {
		if i == j[1] {
			continue
		}
		for m := i; m > j[1]; m-- {
			people[m] = people[m-1]
		}
		people[j[1]] = j
	}
	return people
}
