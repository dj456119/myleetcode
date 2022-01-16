/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2022-01-08 16:30:10
 * @LastEditors: cm.d
 * @LastEditTime: 2022-01-08 16:44:21
 */
package myleetcode

import (
	"fmt"
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	result := [][]int{}
	sort.Slice(nums, func(x, y int) bool {
		return nums[x] < nums[y]
	})
	rmap := make(map[string]bool)
	for i := range nums {
		for j := range nums {
			if j <= i {
				continue
			}
			m := j + 1
			n := len(nums) - 1
			for {
				if m >= n {
					break
				}
				tt := nums[i] + nums[j] + nums[m] + nums[n]
				if tt == target {

					key := fmt.Sprintf("%d_%d_%d_%d", nums[i], nums[j], nums[m], nums[n])
					if _, ok := rmap[key]; ok {

					} else {
						result = append(result, []int{nums[i], nums[j], nums[m], nums[n]})
						rmap[key] = true
					}
				}
				if tt > target {
					n--
				} else {
					m++
				}

			}
		}
	}
	return result
}
