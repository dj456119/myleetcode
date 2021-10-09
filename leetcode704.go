/**
二分查找
*/
package myleetcode

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	if len(nums) == 1 {
		if nums[0] == target {
			return 0
		} else {
			return -1
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
		if end == start || end == start+1 {
			if nums[start] == target {
				return start
			} else if nums[end] == target {
				return end
			} else {
				return -1
			}
		}
	}

}
func Search(nums []int, target int) int {
	return search(nums, target)
}
