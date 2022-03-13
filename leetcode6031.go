/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2022-03-13 10:31:33
 * @LastEditors: cm.d
 * @LastEditTime: 2022-03-13 11:37:03
 */
package myleetcode

import (
	"math"
)

func findKDistantIndices(nums []int, key int, k int) []int {
	z := []int{}
	for i := range nums {
		if nums[i] == key {
			z = append(z, i)
		}
	}
	result := []int{}
	for i := range nums {
		for j := range z {
			if int(math.Abs(float64(i-z[j]))) <= k {
				result = append(result, i)
				break
			}
		}

	}
	return result
}
