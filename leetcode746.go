/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2022-01-16 01:03:11
 * @LastEditors: cm.d
 * @LastEditTime: 2022-01-16 01:42:59
 */
package myleetcode

func minCostClimbingStairs(cost []int) int {
	if len(cost) == 0 || len(cost) == 1 {
		return 0
	}
	fn2 := 0
	fn1 := min(cost[0], cost[1])
	now := 2
	fn := 0

	if now == len(cost) {
		return fn1
	}
	for now != len(cost) {
		fn = min(fn1+cost[now], fn2+cost[now-1])
		fn1, fn2 = fn, fn1
		now++
	}
	return fn
}

func min(fn1, fn2 int) int {
	if fn1 < fn2 {
		return fn1
	}
	return fn2
}
