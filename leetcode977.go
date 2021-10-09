package myleetcode

func sortedSquares(nums []int) []int {
	if len(nums) == 0 {
		return nums
	}
	if len(nums) == 1 {
		return []int{nums[0] * nums[0]}
	}
	result := make([]int, len(nums))
	start := 0
	end := len(nums) - 1
	i := end
	for {
		start2 := nums[start] * nums[start]
		end2 := nums[end] * nums[end]
		if end2 >= start2 {
			result[i] = end2
			end--
		} else {
			result[i] = start2
			start++
		}
		i--
		if start == end {
			result[0] = nums[start] * nums[start]
			return result
		}
	}

}

func SortedSquares(nums []int) []int {
	return sortedSquares(nums)
}
