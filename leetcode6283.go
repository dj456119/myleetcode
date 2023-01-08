package myleetcode

func maximumCount(nums []int) int {
	m1 := 0
	m2 := 0
	for i := range nums {
		if nums[i] == 0 {
			continue
		}
		if nums[i] > 0 {
			m1++
		} else {
			m2++
		}
	}
	if m1 >= m2 {
		return m1
	}
	return m2
}
