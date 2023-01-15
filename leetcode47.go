package myleetcode

import "sort"

func permuteUnique(nums []int) [][]int {
	sort.Slice(nums, func(i, j int) bool { return nums[i] < nums[j] })
	return backtrace47(nums, 0, make([]bool, len(nums)), []int{})
}

func backtrace47(nums []int, pos int, visited []bool, result []int) [][]int {
	if pos == len(nums) {
		return [][]int{append([]int{}, result...)}
	}
	rrr := [][]int{}
	for i := range nums {
		if visited[i] || i != 0 && nums[i] == nums[i-1] && !visited[i-1] {
			continue
		}
		visited[i] = true
		result = append(result, nums[i])
		r := backtrace47(nums, pos+1, visited, result)
		visited[i] = false
		result = result[:len(result)-1]
		rrr = append(rrr, r...)
	}
	return rrr
}
