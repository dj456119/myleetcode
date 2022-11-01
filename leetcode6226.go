package myleetcode

func destroyTargets(nums []int, space int) int {
	z := make(map[int]int)
	z1 := make(map[int][]int)
	for i := range nums {
		z[nums[i]]++
	}
	for i := range nums {
		c := nums[i] % space
		if f, ok := z1[c]; ok {
			f = append(f, nums[i])
			z1[c] = f
		} else {
			f = []int{nums[i]}
			z1[c] = f
		}
	}
	max := 0
	maxR := []int{}
	for _, v := range z1 {
		if len(v) > max {
			max = len(v)
		}
	}
	for _, v := range z1 {
		if len(v) == max {
			maxR = append(maxR, v...)
		}
	}
	result := 100000000000
	for i := range maxR {
		if maxR[i] < result {
			result = maxR[i]
		}
	}
	//    fmt.Println(maxR, max)
	return result
}
