package myleetcode

func reverseWords(s string) string {
	bytes := []byte(s)
	spaceB := []byte(" ")[0]
	var prev int
	for i, byte := range bytes {
		if byte == spaceB {
			reverseString557(bytes, prev, i-1)
			prev = i + 1
		}
	}
	if prev < len(s)-1 {
		reverseString557(bytes, prev, len(bytes)-1)
	}
	return string(bytes)
}

func reverseString557(s []byte, start, end int) {
	if len(s) == 0 || len(s) == 1 {
		return
	}
	if start >= end {
		return
	}
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
func ReverseWords(s string) string {
	return reverseWords(s)
}
