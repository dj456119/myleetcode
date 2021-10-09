package myleetcode

func searchInsert(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		if nums[0] == target {
			return 0
		} else if nums[0] > target {
			return 0
		} else {
			return 1
		}
	}
	start := 0
	end := len(nums) - 1
	for {
		middle := (end + start) / 2
		if nums[middle] > target {
			end = middle
		} else if nums[middle] < target {
			start = middle
		} else {
			return middle
		}

		if end == start+1 {
			if nums[end] == target {
				return end
			} else if nums[end] < target {
				return end + 1
			}
		}

		if end == start || end == start+1 {
			if nums[start] == target {
				return start
			} else if nums[start] > target {
				return start
			} else {
				return start + 1
			}
		}
	}

}

func SearchInsert(nums []int, target int) int {
	return searchInsert(nums, target)
}
