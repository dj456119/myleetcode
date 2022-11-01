package myleetcode

func countSubarrays(nums []int, minK int, maxK int) int64 {
	var result int64
	for i := range nums {
		if nums[i] < minK || nums[i] > maxK {
			continue
		}
		flagMin := false
		flagMax := false
		for j := i; j < len(nums); j++ {
			if nums[j] < minK || nums[j] > maxK {
				continue
			}
			if nums[j] == minK {
				flagMin = true
			}
			if nums[j] == maxK {
				flagMax = true
			}
			if flagMax && flagMin {
				result++
			}
		}
	}
	return result
}
