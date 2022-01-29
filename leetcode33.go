/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2022-01-29 22:15:36
 * @LastEditors: cm.d
 * @LastEditTime: 2022-01-29 22:59:51
 */
package myleetcode

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	start := 0
	end := len(nums) - 1
	mid := (start + end) / 2
	z := -1
	for ; nums[mid] != target; mid = (start + end) / 2 {
		if start >= end {
			return -1
		}
		if nums[start] <= nums[mid] {
			z = bsearch(nums, target, start, mid)
			start = mid + 1
		} else if nums[mid] <= nums[end] {
			z = bsearch(nums, target, mid+1, end)
			end = mid
		}
		if z != -1 {
			return z
		}
	}
	return mid
}

func bsearch(nums []int, target int, start, end int) int {
	if len(nums) == 0 {
		return -1
	}
	mid := (start + end) / 2
	for ; nums[mid] != target; mid = (start + end) / 2 {
		if start >= end {
			return -1
		}
		if nums[mid] > target {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	return mid
}
