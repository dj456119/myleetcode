/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2022-01-02 11:36:03
 * @LastEditors: cm.d
 * @LastEditTime: 2022-01-02 11:42:18
 */

package myleetcode

import (
	"sort"
)

func asteroidsDestroyed(mass int, asteroids []int) bool {
	sort.Slice(asteroids, func(i, j int) bool {
		return asteroids[i] <= asteroids[j]
	})

	for _, j := range asteroids {
		if j < mass {
			mass = mass + j
		} else {
			return false
		}
	}
	return true
}
