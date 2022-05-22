package myleetcode

func percentageLetter(s string, letter byte) int {
	acount := 0
	for i := range s {
		if s[i] == letter {
			acount++
		}
	}
	return acount * 100 / len(s)
}
