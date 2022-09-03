package myleetcode

import "sort"

func answerQueries(nums []int, queries []int) []int {
	sort.Slice(nums, func(i, j int) bool { return nums[i] < nums[j] })
	result := []int{}
	for i := range queries {
		r := 0
		t := queries[i]
		for j := range nums {
			if t >= nums[j] {
				t = t - nums[j]
				r++
			}
		}
		result = append(result, r)
	}
	return result
}
