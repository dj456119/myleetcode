package myleetcode

func closetTarget(words []string, target string, startIndex int) int {
	z := []string{}
	for i := 0; i < 3; i++ {
		z = append(z, words...)
	}
	flag := false
	for i := range words {
		if words[i] == target {
			flag = true
		}
	}
	if !flag {
		return -1
	}
	left := 0
	right := 0
	for i := startIndex + len(words); ; i++ {
		if z[i] == target {
			right = i - (startIndex + len(words))
			break
		}
	}
	for i := startIndex + len(words); ; i-- {
		if z[i] == target {
			left = startIndex + len(words) - i
			break
		}
	}
	if left > right {
		return right
	}
	return left
}
