/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2022-02-01 19:18:39
 * @LastEditors: cm.d
 * @LastEditTime: 2022-02-02 00:42:26
 */
package myleetcode

func coinChange(coins []int, amount int) int {
	dparr := make([]int, amount+1)
	INIT := 999999
	for x := range dparr {
		if x == 0 {
			dparr[0] = 0
			continue
		}
		dparr[x] = INIT
		for _, i := range coins {
			if x >= i {
				dparr[x] = min(dparr[x], dparr[x-i]+1)
			}
		}

	}

	if dparr[len(dparr)-1] == INIT {
		return -1
	}
	return dparr[len(dparr)-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
