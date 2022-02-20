/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2022-02-19 22:33:59
 * @LastEditors: cm.d
 * @LastEditTime: 2022-02-19 22:37:30
 */
package myleetcode

func countPairs(nums []int, k int) int {
	result := 0
	for i := range nums {
		for j := range nums {
			if j <= i {
				continue
			}
			if nums[i] == nums[j] && i*j%k == 0 {
				result++
			}
		}
	}
	return result
}
