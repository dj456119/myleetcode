package myleetcode

func maximumsSplicedArray(nums1 []int, nums2 []int) int {
	nums3 := []int{}
	sum1 := 0
	for i := range nums1 {
		sum1 = nums1[i] + sum1
	}
	sum2 := 0
	for i := range nums2 {
		sum2 = sum2 + nums2[i]
	}
	for i := range nums1 {
		nums3 = append(nums3, nums1[i]-nums2[i])
	}
	c := []int{}
	sum := 0
	for i := range nums3 {
		if i == 0 {
			sum = i
			continue
		}
		if nums3[i-1] < 0 && nums3[i] < 0 {
			sum = sum + nums3[i]
			continue
		} else if nums3[i] < 0 {
			c = append(c, sum)
			sum = nums3[i]
			continue
		}
		if nums3[i] < 0 {
			c = append(c, sum)
			sum = nums3[i]
			continue
		}
		sum = sum + nums3[i]
	}
	result := 0
	if sum1 >= sum2 {
		result = sum1
	} else {
		result = sum2
	}
	for i := range c {
		if c[i] < 0 {
			if -c[i]+sum2 > result {
				result = -c[i] + sum2
			}
		} else {
			if c[i]+sum1 > result {
				result = c[i] + sum1
			}
		}
	}
	return result
}
