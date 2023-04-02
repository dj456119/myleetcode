package myleetcode

func findTheLongestBalancedSubstring(s string) int {
	max := 0
	for i := 0; i < len(s); i++ {
		for j := i + 1; j <= len(s); j++ {
			z := howLength([]byte(s[i:j]))
			//fmt.Println(z, i, j, s[i:j])
			if z > max {
				max = z
			}
		}
	}
	return max
}

func howLength(b []byte) int {
	count0 := 0
	count1 := 0
	for i := range b {
		if i == 0 {
			if b[i] == '1' {
				return 0
			} else {
				count0++
			}
			continue
		}
		if b[i-1] == '1' && b[i] == '0' {
			return 0
		}
		if b[i] == '0' {
			count0++
		}
		if b[i] == '1' {
			count1++
		}
	}
	if count0 != count1 {
		return 0
	}
	return count0 + count1
}
