package myleetcode

func rotate(nums []int, k int) {
	if len(nums) == 0 || len(nums) == 1 {
		return
	}
	k = k % len(nums)
	if k == 0 {
		return
	}

	reserve(nums, 0, len(nums)-1)
	reserve(nums, 0, k-1)
	reserve(nums, k, len(nums)-1)
}

func reserve(nums []int, start, end int) {
	if end == start {
		return
	}
	for {
		nums[end], nums[start] = nums[start], nums[end]
		if end-start == 1 {
			return
		}
		start++
		end--
		if start == end {
			return
		}
	}
}

func Rotate(nums []int, k int) {
	rotate(nums, k)
}
