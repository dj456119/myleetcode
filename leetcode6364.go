package myleetcode

import "sort"

func miceAndCheese(reward1 []int, reward2 []int, k int) int {
	z := make([][]int, len(reward1))
	for i := range z {
		z[i] = []int{reward1[i], reward2[i]}
	}
	sort.Slice(
		z, func(i, j int) bool {
			if z[i][0] >= z[i][1] {
				if z[j][0] < z[j][1] {
					return true
				} else {
					if z[i][0]-z[i][1] >= z[j][0]-z[j][1] {
						return true
					} else {
						return false
					}
				}
			} else {
				if z[j][0] > z[j][1] {
					return false
				} else {
					if z[i][1]-z[i][0] >= z[j][1]-z[j][0] {
						return true
					} else {
						return false
					}
				}
			}
			return false
		})
	sum := 0
	for i := range z {
		if i < k {
			sum = sum + z[i][0]
		} else {
			sum = sum + z[i][1]
		}
	}
	return sum
}
