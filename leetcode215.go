/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2022-01-29 09:49:25
 * @LastEditors: cm.d
 * @LastEditTime: 2022-01-29 09:59:08
 */
package myleetcode

func findKthLargest(nums []int, k int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	start := 0
	end := len(nums) - 1
	pos := partition(nums, start, end)
	for pos != k {
		if pos > k {
			end = pos
		} else {
			start = pos
		}
		pos = partition(nums, start, end)
	}
	return pos
}

func partition(nums []int, start, end int) int {
	mid := (start + end) / 2
	for {
		for nums[start] < nums[mid] && start < mid {
			start++
		}
		nums[start], nums[mid] = nums[mid], nums[start]
		mid = start
		for nums[end] > nums[mid] && end > mid {
			end--
		}
		nums[end], nums[mid] = nums[mid], nums[end]
		mid = end
		if start >= end {
			return mid
		}
	}

}
