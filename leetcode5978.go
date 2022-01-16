/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2022-01-09 10:45:47
 * @LastEditors: cm.d
 * @LastEditTime: 2022-01-09 23:58:56
 */
package myleetcode

import "fmt"

func wordCount(startWords []string, targetWords []string) int {
	ss := make([]int, len(startWords))
	ts := make([]int, len(targetWords))
	fmt.Println(len(startWords))
	resultMap := make(map[string]bool)
	for i, j := range startWords {
		bm := 0
		for _, n := range j {
			bm = bm | (1 << (n - 'a'))
		}
		ss[i] = bm
	}
	for i, j := range targetWords {
		bm := 0
		for _, n := range j {
			bm = bm | (1 << (n - 'a'))
		}
		ts[i] = bm
	}
	for i, j := range ss {
		for m, n := range ts {
			z := hammingDistance(j, n)

			if z == 1 && (j&n == j) {
				fmt.Println(startWords[i], targetWords[m])
				if _, ok := resultMap[targetWords[m]]; !ok {
					resultMap[targetWords[m]] = true
				}
			}
		}
	}
	return len(resultMap)
}

func hammingDistance(x, y int) (ans int) {
	for s := x ^ y; s > 0; s &= s - 1 {
		ans++
	}
	return
}
