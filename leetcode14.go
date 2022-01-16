/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2022-01-08 14:31:06
 * @LastEditors: cm.d
 * @LastEditTime: 2022-01-08 14:40:29
 */
package myleetcode

func longestCommonPrefix(strs []string) string {
	minL := -1
	for _, j := range strs {

		if len(j) < minL || minL == -1 {
			minL = len(j)
		}
	}
	if minL == 0 {
		return ""
	}
	for m := 0; m < minL; m++ {
		prev := strs[0][m]
		for _, j := range strs {
			if j[m] != prev {
				return j[:m]
			}
		}
	}
	return strs[0][:minL]
}
