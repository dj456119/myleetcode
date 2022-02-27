/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2022-01-16 20:55:18
 * @LastEditors: cm.d
 * @LastEditTime: 2022-02-25 22:15:24
 */
package myleetcode

import "sort"

func combinationSum2(candidates []int, target int) [][]int {
	sort.Slice(candidates, func(x, y int) bool {
		return candidates[x] < candidates[y]
	})
	return backtrace40(candidates, 0, target, []int{})
}

func backtrace40(nums []int, pos int, target int, result []int) [][]int {
	if target < 0 {
		return nil
	}
	if target == 0 {
		rr1 := make([]int, len(result))
		for i := range result {
			rr1[i] = nums[result[i]]
		}
		return [][]int{rr1}
	}

	if len(nums) == pos {
		return nil
	}
	if nums[pos] > target {
		return nil
	}
	rr := [][]int{}
	length := len(result)
	ok := (pos != 0) && ((len(result) == 0 && nums[pos] == nums[pos-1]) || (len(result) != 0 && nums[pos] == nums[pos-1] && result[len(result)-1] != pos-1))
	if !ok {
		result = append(result, pos)
		r1 := backtrace40(nums, pos+1, target-nums[pos], result)
		if r1 != nil {
			rr = append(rr, r1...)
		}
		result = result[:length]
	}
	r2 := backtrace40(nums, pos+1, target, result)
	if r2 != nil {
		rr = append(rr, r2...)
	}
	return rr
}
