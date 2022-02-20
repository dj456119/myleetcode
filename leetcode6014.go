/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2022-02-20 11:26:08
 * @LastEditors: cm.d
 * @LastEditTime: 2022-02-20 11:34:23
 */
package myleetcode

import "sort"

func repeatLimitedString(s string, repeatLimit int) string {
	z := []byte(s)
	sort.Slice(z, func(i, j int) bool {
		if z[i] < z[j] {
			return true
		}
		return false
	})
	i := 0
	record := z[0]
	count := 0
	for {
		if i == len(s) {
			break
		}
		if z[i] == record {
			count++
			if count > repeatLimit {
				j := i
				for z[j] == record {
					j++
					if j == len(s) {
						return string(z[:i+1])
					}
				}
				z[j], z[i] = z[i], z[j]
				continue
			}
		} else {
			record = z[i]
			count = 1
		}
		i++
	}
	return string(z)
}
