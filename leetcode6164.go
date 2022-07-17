package myleetcode

func maximumSum(nums []int) int {
	cm := make(map[int][]int)
	for i := range nums {
		z := nums[i]
		sm := 0
		for z != 0 {
			sm = sm + z%10
			z = z / 10
		}
		if ll, ok := cm[sm]; ok {
			ll = append(ll, nums[i])
			cm[sm] = ll
		} else {
			ll = make([]int, 0)
			ll = append(ll, nums[i])
			cm[sm] = ll
		}
	}
	//  fmt.Println(cm)
	max := -1
	for _, v := range cm {
		if len(v) > 1 {
			max1 := 0
			max2 := 0
			for i := range v {
				if v[i] > max1 {
					max2 = max1
					max1 = v[i]
					continue
				}
				if v[i] > max2 && v[i] <= max1 {
					max2 = v[i]
				}
			}
			if max1+max2 > max {
				max = max1 + max2
			}
		}
	}
	return max
}
