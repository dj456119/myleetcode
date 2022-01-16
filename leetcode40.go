/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2022-01-16 20:55:18
 * @LastEditors: cm.d
 * @LastEditTime: 2022-01-16 21:13:34
 */
package myleetcode

import (
	"fmt"
	"sort"
)

func combinationSum2(candidates []int, target int) [][]int {
	rr := make(map[string]bool)
	return digui(candidates, target, 0, []int{}, rr)
}

func digui(candidates []int, target int, now int, result []int, rr map[string]bool) [][]int {
	if now == len(candidates) {
		if target == 0 {
			sort.Slice(result, func(i, j int) bool { return result[i] < result[j] })
			key := ""
			for _, j := range result {
				key = fmt.Sprintf("_%d", j)
			}
			if _, ok := rr[key]; ok {
				return nil
			}
			rr[key] = true
			return [][]int{result}
		}
		return nil
	}
	fin := [][]int{}
	z1 := digui(candidates, target, now+1, result, rr)
	r := make([]int, len(result))
	copy(r, result)
	target = target - candidates[now]
	r = append(r, candidates[now])
	z2 := digui(candidates, target, now+1, r, rr)
	if z1 != nil {
		fin = append(fin, z1...)
	}
	if z2 != nil {
		fin = append(fin, z2...)
	}
	return fin
}
