/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2022-01-08 14:43:33
 * @LastEditors: cm.d
 * @LastEditTime: 2022-01-08 15:01:46
 */
package myleetcode

import (
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	sort.Slice(nums, func(x, y int) bool {
		return nums[x] < nums[y]
	})
	result := -1
	r := 0
	for i, j := range nums {
		m := 0
		n := len(nums) - 1
		for {
			if m == i {
				m++
			}
			if n == i {
				n--
			}
			if m >= n {
				break
			}
			tt := nums[m] + nums[n] + j
			if tt > target {
				n--
			} else {
				if tt < target {
					m++
				} else {
					return tt
				}
			}
			if math.Abs(float64(tt-target)) < float64(result) || result == -1 {
				result = int(math.Abs(float64(tt - target)))
				r = tt
			}
		}
	}
	return r
}
