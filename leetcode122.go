/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2022-02-12 16:51:13
 * @LastEditors: cm.d
 * @LastEditTime: 2022-02-12 17:28:45
 */
package myleetcode

func maxProfit(prices []int) int {
	dparr := make([][]int, len(prices))
	for i := range dparr {
		dparr[i] = make([]int, 2)
	}
	for i := range prices {
		if i == 0 {
			dparr[0][0] = 0
			dparr[0][1] = -prices[i]
			continue
		}
		dparr[i][0] = max(dparr[i-1][0], dparr[i-1][1]+prices[i])
		dparr[i][1] = max(dparr[i-1][1], dparr[i-1][0]-prices[i])
	}
	return dparr[len(dparr)-1][0]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
