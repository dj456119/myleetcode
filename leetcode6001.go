/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2022-02-06 10:38:44
 * @LastEditors: cm.d
 * @LastEditTime: 2022-02-06 10:51:30
 */
package myleetcode

import (
	"math"
	"sort"
)

func smallestNumber(num int64) int64 {
	z := []int{}
	i := 0
	zeroCount := 0
	for num != 0 {
		m := num % 10
		if num < 0 {
			z[i] = int(m)
			i++
		} else {
			if m == 0 {
				zeroCount = 0
			}
		}
		num = num / 10
	}
	if num < 0 {
		sort.Slice(z, func(i, j int) bool {
			return z[i] > z[j]
		})
		var sum int64 = 0
		for i := range z {
			sum = sum*10 + int64(z[i])
		}
		return -sum
	} else {
		sort.Slice(z, func(i, j int) bool {
			return z[i] < z[j]
		})
		var sum int64 = 0
		sum = int64(z[0]) * int64(math.Pow(10, float64(zeroCount)))
		for i := range z {
			if i != 0 {
				sum = sum*10 + int64(z[i])
			}
		}
		return sum
	}
}
