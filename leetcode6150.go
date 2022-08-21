package myleetcode

func smallestNumber6150(pattern string) string {
	rr := ""
	for i := 1; i <= 9; i++ {
		used := make([]int, 10)
		used[i] = 1
		result := []byte{}
		result = append(result, byte('0'+i))
		rr1 := backtrace6150(pattern, 0, i, used, result)
		if rr1 == "" {
			rr = rr1
		}
		if rr1 != "" {
			if rr1 < rr {
				rr = rr1
			}
		}
	}
	return rr
}

func backtrace6150(pattern string, pos int, prev int, used []int, result []byte) string {
	if pos == len(pattern) {
		return string(result)
	}
	rr := ""
	if pattern[pos] == 'I' {
		for i := 1; i <= 9; i++ {
			if used[i] == 1 {
				continue
			}
			if i < prev {
				continue
			}
			r1 := make([]byte, len(result))
			u1 := make([]int, len(used))
			copy(r1, result)
			copy(u1, used)
			u1[i] = 1
			r1 = append(result, byte(i+'0'))
			rr1 := backtrace6150(pattern, pos+1, i, u1, r1)
			if rr1 == "" {
				rr = rr1
			}
			if rr1 != "" {
				if rr1 < rr {
					rr = rr1
				}
			}
		}
	} else {
		for i := 1; i <= 9; i++ {
			if used[i] == 1 {
				continue
			}
			if i > prev {
				continue
			}
			r1 := make([]byte, len(result))
			u1 := make([]int, len(used))
			copy(r1, result)
			copy(u1, used)
			u1[i] = 1
			r1 = append(result, byte(i+'0'))
			rr1 := backtrace6150(pattern, pos+1, i, u1, r1)
			if rr1 == "" {
				rr = rr1
			}
			if rr1 != "" {
				if rr1 < rr {
					rr = rr1
				}
			}
		}
	}
	return rr
}
