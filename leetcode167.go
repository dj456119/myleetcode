package myleetcode

func twoSum167(numbers []int, target int) []int {
	if len(numbers) == 0 || len(numbers) == 1 {
		return nil
	}
	start := 0
	end := len(numbers) - 1
	for {
		sum := numbers[start] + numbers[end]
		if sum > target {
			end--
		} else if sum < target {
			start++
		} else {
			return []int{start + 1, end + 1}
		}
		if start >= end {
			return nil
		}
	}
}

func TwoSum167(numbers []int, target int) []int {
	return twoSum167(numbers, target)
}
