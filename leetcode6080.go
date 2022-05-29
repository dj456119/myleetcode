package myleetcode

func totalSteps(nums []int) int {
	c := make([]int, len(nums))
	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i] > nums[i+1] {
			c[i] = c[i+1] + 1
		} else {
			c[i] = 0
		}
	}
	c1 := make([]int, len(nums))
	m := -1
	c2 := []int{}
	for i := range nums {
		if nums[i] >= m {
			m = nums[i]
			c2 = append(c2, i)
		}
	}
	for i := range c2 {
		if i == len(c2)-1 {
			c1[c2[i]] = len(c2) - c2[i] - 1
		}
		c1[c2[i]] = c2[i+1] - c2[i] - 1
	}
	max := 0
	for i := range c {
		if c[i] > max {
			max = c[i]
		}
	}
	return max
}
