package myleetcode

func countDistinctIntegers(nums []int) int {
	c := make(map[int64]bool)
	for i := range nums {
		c[int64(nums[i])] = true
		d := resNum(nums[i])
		c[d] = true
	}
	return len(c)
}

func resNum(n int) int64 {
	var result int64
	for n != 0 {
		z := n % 10
		result = result*10 + int64(z)
		n = n / 10
	}
	return result
}
