package main

func longestAlternatingSubarray(nums []int, threshold int) int {
	max := 0
	for i := range nums {
		if nums[i]%2 != 0 {
			continue
		}

		for j := range nums {
			flag := false
			if j < i {
				continue
			}
			z := nums[i : j+1]

			for m := range z {
				if z[m] > threshold {
					flag = true
				}
				if m == len(z)-1 {
					break
				}
				if z[m]%2 == z[m+1]%2 {
					flag = true
				}

			}
			if flag {
				continue
			}
			//     fmt.Println(z)
			if len(z) > max {
				max = len(z)
			}
		}
	}
	return max
}
