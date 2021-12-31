/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2021-12-31 11:57:49
 * @LastEditors: cm.d
 * @LastEditTime: 2021-12-31 12:10:18
 */
package myleetcode

import (
	"math"
)

func mySqrt(x int) int {
	j := x
	prev := j
	for {
		if math.Pow(float64(j), 2) > float64(x) {
			prev = j
			j = j / 2

		} else {

			if math.Pow(float64(j), 2) == float64(x) {
				return j
			}
			break
		}
	}

	for i := j; i <= prev; i++ {
		if math.Pow(float64(i), 2) > float64(x) {
			return i - 1
		}
		if math.Pow(float64(i), 2) == float64(x) {
			return i
		}
	}
	return 0
}
