package myleetcode

func findDifference(nums1 []int, nums2 []int) [][]int {
	m1 := make(map[int]int)
	m2 := make(map[int]int)
	for i := range nums1 {
		m1[nums1[i]] = 1
	}
	for i := range nums2 {
		m2[nums2[i]] = 1
	}
	result := [][]int{}
	r1 := []int{}
	r2 := []int{}
	for k := range m1 {
		if m2[k] != 1 {
			r1 = append(r1, k)
		}
	}
	for k := range m2 {
		if m1[k] != 1 {
			r2 = append(r2, k)
		}
	}
	result = append(result, r1)
	result = append(result, r2)
	return result
}
