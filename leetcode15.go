/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2022-01-01 00:38:47
 * @LastEditors: cm.d
 * @LastEditTime: 2022-01-01 00:56:39
 */
package myleetcode

import (
	"fmt"
	"sort"
)

var tmap map[string][]int

func threeSum(nums []int) [][]int {
	tmap = make(map[string][]int)
	ncount(nums, 0, 3, []int{})
	result := [][]int{}
	for _, j := range tmap {
		sum := 0
		for _, m := range j {
			sum = sum + m
		}
		if sum == 0 {
			result = append(result, j)
		}
	}
	return result
}

func ncount(nums []int, now, all int, result []int) [][]int {
	if len(result) == all {
		sort.Ints(result)
		key := fmt.Sprintf("%d_%d_%d", result[0], result[1], result[2])
		if _, ok := tmap[key]; !ok {
			tmap[key] = result
		}
		return [][]int{result}
	}
	if now == len(nums) {
		return [][]int{}
	}
	r1 := make([]int, len(result)+1)
	copy(r1, result)
	r1[len(result)] = nums[now]
	rr1 := ncount(nums, now+1, all, result)
	rr2 := ncount(nums, now+1, all, r1)
	rrr := [][]int{}
	if len(rr1) != 0 {
		rrr = append(rrr, rr1...)
	}
	if len(rr2) != 0 {
		rrr = append(rrr, rr2...)
	}
	return rrr
}
