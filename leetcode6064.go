package myleetcode

import "sort"

func maxConsecutive(bottom int, top int, special []int) int {
	special = append(special, top+1)
	special = append(special, bottom-1)
	sort.Slice(special, func(a, b int) bool {
		return special[a] < special[b]
	})
	max := 0
	for i := range special {
		if i == 0 {
			continue
		}
		if special[i]-special[i-1]-1 > max {
			max = special[i] - special[i-1] - 1
		}
	}
	return max
}
