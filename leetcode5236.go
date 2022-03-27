package myleetcode

func minDeletion(nums []int) int {
	result := 0
	pos := 0
	sj := 0
	for {
		if pos == len(nums)-1 {
			sj++
			break
		}
		if sj%2 == 0 && nums[pos] == nums[pos+1] {
			result++
			pos++
			continue
		}
		sj++
		pos++
	}
	if sj%2 == 1 {
		result++
	}
	return result
}
