package myleetcode

func reverseString(s []byte) {
	if len(s) == 0 || len(s) == 1 {
		return
	}
	start := 0
	end := len(s) - 1
	for {
		s[start], s[end] = s[end], s[start]
		if start+1 == end {
			return
		}
		start++
		end--
		if start == end {
			return
		}
	}
}

func ReverseString(s []byte) {
	reverseString(s)
}
