/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2022-02-14 00:40:46
 * @LastEditors: cm.d
 * @LastEditTime: 2022-02-14 01:36:15
 */
package myleetcode

func maxSlidingWindow(nums []int, k int) []int {
	indexQueue := []int{}
	for i := 0; i <= k-1; i++ {
		for len(indexQueue) != 0 && nums[i] > nums[indexQueue[len(indexQueue)-1]] {
			indexQueue = indexQueue[:len(indexQueue)-1]
		}
		indexQueue = append(indexQueue, i)
	}
	result := []int{}
	result = append(result, nums[indexQueue[0]])

	for i := k; i < len(nums); i++ {
		c := 0
		for z := range indexQueue {
			if indexQueue[z] <= i-k {
				c++
			} else {
				break
			}
		}
		indexQueue = indexQueue[c:]
		for len(indexQueue) != 0 && nums[i] > nums[indexQueue[len(indexQueue)-1]] {
			indexQueue = indexQueue[:len(indexQueue)-1]
		}
		indexQueue = append(indexQueue, i)
		result = append(result, nums[indexQueue[0]])
	}
	return result
}
