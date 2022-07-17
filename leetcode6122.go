package myleetcode

import "sort"

func minOperations(nums []int, numsDivide []int) int {
	sort.Slice(nums, func(i, j int) bool { return nums[i] < nums[j] })
	min := 9999999999
	for i := range numsDivide {
		if numsDivide[i] < min {
			min = numsDivide[i]
		}
	}

	for i := range nums {
		if nums[i] > min {
			return -1
		}
		flag := true
		for j := range numsDivide {
			if numsDivide[j]%nums[i] != 0 {
				flag = false
				break
			}
		}
		if flag {
			return i
		}
	}
	return -1
}
