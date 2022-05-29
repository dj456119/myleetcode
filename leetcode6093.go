package myleetcode

func digitCount(num string) bool {
	for i := 0; i < len(num); i++ {
		c := int(num[i] - '0')
		sum := 0
		for j := 0; j < len(num); j++ {
			if int(num[j]-'0') == i {
				sum++
			}
		}
		if sum != c {
			return false
		}
	}
	return true
}
