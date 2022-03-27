package myleetcode

import "math"

func kthPalindrome(queries []int, intLength int) []int64 {
	result := []int64{}
	wei := intLength / 2
	if intLength%2 == 0 {
		wei = intLength/2 - 1
	}
	max := 9 * math.Pow10(wei)
	if intLength == 1 {
		for i := range queries {
			if int64(queries[i]) > int64(max) {
				result = append(result, -1)
				continue
			}
			result = append(result, int64(queries[i]))
		}
		return result
	}
	if intLength == 2 {
		for i := range queries {
			if int64(queries[i]) > int64(max) {
				result = append(result, -1)
				continue
			}
			result = append(result, int64(queries[i]*10+queries[i]))
		}
		return result
	}

	for _, j := range queries {
		if int64(j) > int64(max) {
			result = append(result, -1)
			continue
		}

		start := make([]int, intLength)
		start[0] = 1
		start[len(start)-1] = 1
		center := intLength / 2
		if intLength%2 == 0 {
			center = intLength/2 - 1
		}

		var temp int64
		for i := 0; i <= center; i++ {
			temp = temp*10 + int64(start[i])
		}
		//    fmt.Println(j, temp)
		temp = temp + int64(j-1)
		//    fmt.Println(j, temp)
		z := center
		for temp != 0 {
			start[z] = int(temp % 10)
			temp = temp / 10
			z--
		}
		//fmt.Println(start)
		i := 0
		q := len(start) - 1
		for i < q {
			start[q] = start[i]
			i++
			q--
		}

		var rr int64
		for i := range start {
			rr = rr*10 + int64(start[i])
		}
		//  fmt.Println(rr)
		result = append(result, rr)
	}

	return result
}
