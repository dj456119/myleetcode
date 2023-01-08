package myleetcode

func maximumSubarraySum(nums []int, k int) int64 {
	c := make(map[int]int)
	var result int64
	var sum int64
	for i := 0; i < k; i++ {
		c[nums[i]]++
		sum = sum + int64(nums[i])
	}
	if len(c) == k {
		result = sum
	}
	for i := k; i < len(nums); i++ {
		c[nums[i-k]]--
		if c[nums[i-k]] <= 0 {
			delete(c, nums[i-k])
		}
		sum = sum - int64(nums[i-k])
		sum = sum + int64(nums[i])
		c[nums[i]]++
		if len(c) == k && sum > result {
			result = sum
		}
	}
	return result
}
