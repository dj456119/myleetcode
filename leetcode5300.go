/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2022-03-05 23:17:29
 * @LastEditors: cm.d
 * @LastEditTime: 2022-03-05 23:51:32
 */
package myleetcode

import "sort"

func getAncestors(n int, edges [][]int) [][]int {
	reedges := [][]int{}
	for i := range edges {
		b := make([]int, 2)
		b[0] = edges[i][1]
		b[1] = edges[i][0]
		reedges = append(reedges, b)
	}
	zm := make(map[int][]int)
	for i := range reedges {
		start := reedges[i][0]
		end := reedges[i][1]
		if c, ok := zm[start]; ok {
			c = append(c, end)
			zm[start] = c
		} else {
			zm[start] = []int{end}
		}
	}
	result := make([]map[int]bool, n)
	for i := 0; i < n; i++ {
		c, ok := zm[i]
		if !ok {
			result[i] = make(map[int]bool)
			continue
		}
		result[i] = make(map[int]bool)
		queue := make([]int, len(c))
		ready := make([]int, n)
		copy(queue, c)
		for len(queue) != 0 {
			m := queue[0]
			result[i][m] = true
			queue = queue[1:]
			if f, ok := zm[m]; ok && ready[m] == 0 {
				queue = append(queue, f...)
			}
			ready[m] = 1
		}
	}

	r := make([][]int, n)
	for i := range result {
		r[i] = []int{}
		for k := range result[i] {
			r[i] = append(r[i], k)
		}
		sort.Slice(r[i], func(x, y int) bool {
			return r[i][x] < r[i][y]
		})
	}
	return r
}
