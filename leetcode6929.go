package main

import "sort"

func maximumBeauty(nums []int, k int) int {
	sort.Slice(nums, func(x, y int) bool {
		return nums[x] < nums[y]
	})
	result := 0
	for i := range nums {
		q := binarySearch(nums[i:], nums[i]+2*k)
		if q > result {
			result = q
		}
	}
	return result
}
func binarySearch(arr []int, k int) int {
	left, right := 0, len(arr)-1

	for left <= right {
		mid := left + (right-left)/2

		if arr[mid] <= k {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return left
}
