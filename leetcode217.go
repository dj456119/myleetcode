package myleetcode

func containsDuplicate(nums []int) bool {
	type void struct{}
	var temp void
	dataMap := make(map[int]void)
	for _, num := range nums {
		if _, ok := dataMap[num]; ok {
			return true
		}
		dataMap[num] = temp
	}
	return false
}

func ContainsDuplicate(nums []int) bool {
	return containsDuplicate(nums)
}
