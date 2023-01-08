package myleetcode

func subarrayLCM(nums []int, k int) int {
	rr := [][]int{}
	rr = append(rr, []int{})
	for i := range nums {
		if nums[i] > k || k%nums[i] != 0 {
			if len(rr[len(rr)-1]) != 0 {
				rr = append(rr, []int{})
			} else {
				continue
			}
		} else {
			rr[len(rr)-1] = append(rr[len(rr)-1], nums[i])
		}
	}
	for i := range rr {
		z := rr[i]
	}
}
