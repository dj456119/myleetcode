package main

func maximumJumps(nums []int, target int) int {
	max := make([]int, len(nums))
	for i := range nums {
		if i != 0 && max[i] == 0 {
			continue
		}
		for j := i + 1; j < len(nums); j++ {
			if nums[j]-nums[i] <= target && nums[j]-nums[i] >= -target {
				if max[j] > max[i]+1 {
					continue
				}
				max[j] = max[i] + 1
			}
		}
	}
	if max[len(max)-1] == 0 {
		return -1
	}
	return max[len(max)-1]
}
