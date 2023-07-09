package main

func maxNonDecreasingLength(nums1 []int, nums2 []int) int {
	max := 1
	nl1_1 := 1
	nl1_2 := 1
	nl2_1 := 1
	nl2_2 := 1
	for i := range nums1 {
		if i == 0 {
			continue
		}
		if nums1[i] >= nums1[i-1] {
			nl1_1 = nl1_1 + 1
			if nl1_1 > max {
				max = nl1_1
			}
		} else {
			nl1_1 = 1
		}
		if nums1[i] >= nums2[i-1] {
			nl1_2 = nl1_2 + 1
			if nl1_2 > max {
				max = nl1_2
			}
		} else {
			nl1_2 = 1
		}
		if nums2[i] >= nums1[i-1] {
			nl2_1 = nl2_1 + 1
			if nl2_1 > max {
				max = nl2_1
			}
		} else {
			nl2_1 = 1
		}
		if nums2[i] >= nums2[i-1] {
			nl2_2 = nl2_2 + 1
			if nl2_2 > max {
				max = nl2_2
			}
		} else {
			nl2_2 = 1
		}
	}
	return max
}
