package myleetcode

func secondsToRemoveOccurrences(s string) int {
	z := []byte(s)
	result := 0

	for {
		flag := false
		for i := 0; i < len(z)-1; i++ {
			if z[i] == '0' && z[i+1] == '1' {
				flag = true
				z[i], z[i+1] = z[i+1], z[i]
				i++
			}
		}
		if !flag {
			return result
		} else {
			result++
		}
	}

}
