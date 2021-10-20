/*
 * @Descripttion:给你一个长度为 n 的整数数组，每次操作将会使 n - 1 个元素增加 1 。返回让数组所有元素相等的最小操作次数。

 * @version:
 * @Author: cm.d
 * @Date: 2021-10-20 23:02:46
 * @LastEditors: cm.d
 * @LastEditTime: 2021-10-20 23:19:50
 */

package myleetcode

func minMoves(nums []int) int {
	moves := 0
	minPos := 0
	for i, j := range nums {
		if j < nums[minPos] {
			minPos = i
		}
	}
	for i := range nums {
		moves = moves + (nums[i] - nums[minPos])
	}
	return moves
}

func MinMoves(nums []int) int {
	return minMoves(nums)
}
