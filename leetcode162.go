package myleetcode

func findPeakElement(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	start := 0
	end := len(nums) - 1
	for {
		if start == end {
			return nums[start]
		}
		if start+1 == end {
			if nums[start] > nums[end] {
				return start
			}
			return end
		}
		mid := (start + end) / 2
		if nums[mid] > nums[mid+1] && nums[mid] > nums[mid-1] {
			return mid
		}
		if nums[mid] > nums[mid+1] {
			end = mid
			continue
		}
		if nums[mid] > nums[mid-1] {
			start = mid
			continue
		}
		end = mid
	}
}
