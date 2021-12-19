/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2021-12-19 12:45:35
 * @LastEditors: cm.d
 * @LastEditTime: 2021-12-19 12:45:36
 */
package myleetcode

func getDescentPeriods(prices []int) int64 {
	lianxu := 0
	var result int64
	for i := range prices {
		if i != 0 {
			if prices[i-1]-prices[i] == 1 {
				lianxu++
			} else {
				if lianxu != 0 {
					result1 := jisuanlianxu(lianxu)
					result = result + result1
				}
				lianxu = 0
			}
		}
	}
	if lianxu != 0 {
		result1 := jisuanlianxu(lianxu)
		result = result + result1
	}
	return result + int64(len(prices))
}

func jisuanlianxu(lianxu int) int64 {
	return int64(lianxu)*1 + int64(lianxu)*(int64(lianxu)-1)/2
}
