/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2021-10-23 09:43:47
 * @LastEditors: cm.d
 * @LastEditTime: 2022-01-24 02:32:07
 */
/**
二分查找
*/
package myleetcode

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	left := 0
	right := len(nums) - 1
	for {
		if left+1 == right || left == right {
			if nums[left] == target {
				return left
			}
			if nums[right] == target {
				return right
			}
			return -1
		}
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			right = mid
		} else {
			left = mid
		}

	}
}
func Search(nums []int, target int) int {
	return search(nums, target)
}
