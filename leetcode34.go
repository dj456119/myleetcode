/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2022-02-06 21:54:32
 * @LastEditors: cm.d
 * @LastEditTime: 2022-02-06 22:19:32
 */
package myleetcode

func searchRange(nums []int, target int) []int {
	a := bSearchLeft(nums, 0, len(nums)-1, target)
	b := bSearchRight(nums, 0, len(nums)-1, target)
	return []int{a, b}
}

func bSearchLeft(nums []int, start, end int, target int) int {
	mid := (start + end) / 2
	for start <= end {
		if start == end && nums[mid] != target {
			return -1
		}
		if nums[mid] == target {
			if mid > start {
				if nums[mid-1] == target {
					end = mid - 1
					mid = (start + end) / 2
					continue
				}
			}
			return mid
		}
		if nums[mid] > target {
			end = mid
		} else {
			start = mid + 1
		}
		mid = (start + end) / 2
	}
	return -1
}

func bSearchRight(nums []int, start, end int, target int) int {
	mid := (start + end) / 2
	for start <= end {
		if start == end && nums[mid] != target {
			return -1
		}
		if nums[mid] == target {
			if mid < end {
				if nums[mid+1] == target {
					start = mid + 1
					mid = (start + end) / 2
					continue
				}
			}
			return mid
		}
		if nums[mid] > target {
			end = mid
		} else {
			start = mid + 1
		}
		mid = (start + end) / 2
	}
	return -1
}
