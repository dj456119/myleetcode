package main

func continuousSubarrays(nums []int) int64 {
	i := 0
	j := 1
	min := nums[i]
	max := nums[j]
	if nums[i] > nums[j] {
		max = nums[i]
		min = nums[j]
	}

	for {
		if i == len(nums) {
			break
		}

	}
}
