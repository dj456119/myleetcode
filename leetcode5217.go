/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2022-03-05 22:52:09
 * @LastEditors: cm.d
 * @LastEditTime: 2022-03-05 23:08:27
 */
package myleetcode

import "sort"

func sortJumbled(mapping []int, nums []int) []int {
	nnums := [][]int{}
	for i := range nums {
		z := nums[i]
		sum := 0
		if z == 0 {
			sum = mapping[0]
		} else {
			j := 1
			for z != 0 {
				a := z % 10
				z = z / 10
				sum = mapping[a]*j + sum
				j = j * 10
			}
		}

		nnums = append(nnums, []int{sum, nums[i]})
	}
	//fmt.Println(nnums)
	sort.Slice(nnums, func(x, y int) bool {
		if nnums[x][0] < nnums[y][0] {
			return true
		}
		if nnums[x][0] > nnums[y][0] {
			return false
		}
		return x < y
	})
	result := []int{}
	for i := range nnums {
		result = append(result, nnums[i][1])
	}
	return result
}
