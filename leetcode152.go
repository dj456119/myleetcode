/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2022-01-22 19:34:11
 * @LastEditors: cm.d
 * @LastEditTime: 2022-01-23 00:01:53
 */
package myleetcode

import "sort"

func maxProduct(nums []int) int {
	lf := make([]int, len(nums))
	lz := make([]int, len(nums))
	max := nums[0]
	for i, j := range nums {
		if i == 0 {
			lf[0] = j
			lz[0] = j
			continue
		}
		z1 := lf[i-1] * j
		z2 := lz[i-1] * j
		a1, _, a3 := sortz(z1, z2, j)
		lf[i] = a1
		lz[i] = a3
		if lz[i] > max {
			max = lz[i]
		}

	}
	return max
}

func sortz(a1, a2, a3 int) (int, int, int) {
	z := []int{a1, a2, a3}
	sort.Slice(z, func(i, j int) bool { return z[i] < z[j] })
	return z[0], z[1], z[2]
}
