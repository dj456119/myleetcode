/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2022-02-13 23:28:13
 * @LastEditors: cm.d
 * @LastEditTime: 2022-02-21 00:11:32
 */
package myleetcode

func maxCoins(nums []int) int {
	nums = append([]int{1}, nums...)
	nums = append(nums, 1)
	dparr := make([][]int, len(nums))
	for i := range dparr {
		dparr[i] = make([]int, len(nums))
	}
	for left := len(nums) - 1; left >= 0; left-- {
		for right := left + 1; right < len(nums); right++ {
			for i := left + 1; i < right; i++ {
				dparr[left][right] = max312(dparr[left][right], nums[left]*nums[i]*nums[right]+dparr[left][i]+dparr[i][right])
			}
		}
	}
	return dparr[0][len(nums)-1]
}

func max312(a, b int) int {
	if a > b {
		return a
	}
	return b
}
