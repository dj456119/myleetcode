package myleetcode

func moveZeroes(nums []int) {
	if len(nums) == 0 || len(nums) == 1 {
		return
	}
	zeroCount := 0
	for i, num := range nums {
		if num == 0 {
			zeroCount++
		} else {
			nums[i-zeroCount] = num
		}

	}
	for i := len(nums) - zeroCount; i < len(nums); i++ {
		nums[i] = 0
	}
}

func MoveZeroes(nums []int) {
	moveZeroes(nums)
}
