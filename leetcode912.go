/*
 * @Descripttion:给你一个整数数组 nums，请你将该数组升序排列。
 * @version:
 * @Author: cm.d
 * @Date: 2021-10-17 23:43:07
 * @LastEditors: cm.d
 * @LastEditTime: 2021-10-18 00:25:02
 */
package myleetcode

func sortArray(nums []int) []int {
	sort(nums, 0, len(nums)-1)
	return nums
}

func sort(nums []int, start, end int) {
	if end <= start || start >= len(nums) || end >= len(nums) {
		return
	}
	if end == start+1 {
		if nums[end] < nums[start] {
			nums[end], nums[start] = nums[start], nums[end]
		}
		return
	}

	newPos := partion(nums, start, end)
	sort(nums, start, newPos)
	sort(nums, newPos+1, end)
}

func partion(nums []int, start int, end int) int {
	middle := (start + end) / 2
	tempStart := start
	tempEnd := end

	for {
		for nums[middle] >= nums[tempStart] && tempStart < middle {
			tempStart++
		}
		if nums[middle] <= nums[tempStart] {
			nums[middle], nums[tempStart] = nums[tempStart], nums[middle]
			middle = tempStart
		}
		for nums[middle] <= nums[tempEnd] && tempEnd > middle {
			tempEnd--
		}
		if nums[middle] >= nums[tempEnd] {
			nums[middle], nums[tempEnd] = nums[tempEnd], nums[middle]
			middle = tempEnd
		}
		if tempEnd <= tempStart {
			break
		}
	}

	return middle
}

func SortArray(nums []int) []int {
	return sortArray(nums)
}
