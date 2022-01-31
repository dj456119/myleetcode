/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2022-02-01 01:14:21
 * @LastEditors: cm.d
 * @LastEditTime: 2022-02-01 01:26:30
 */
package myleetcode

import "math"

func numSquares(n int) int {
	dparr := []int{}
	dparr = append(dparr, 0)
	for i := 1; i <= n; i++ {
		if i == 1 {
			dparr = append(dparr, 1)
			continue
		}
		z := math.Sqrt(float64(i))
		minz := 0
		for j := 1; float64(j) <= z; j++ {
			z1 := i - j*j
			x := 1 + dparr[z1]
			if j == 1 {
				minz = x
			} else {
				minz = min(x, minz)
			}
		}
		dparr = append(dparr, minz)

	}
	return dparr[n]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
