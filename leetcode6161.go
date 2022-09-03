package myleetcode

func removeStars(s string) string {
	stack := []byte{}
	for i := range s {
		if s[i] == '*' {
			if len(stack) == 0 {
				stack = append(stack, s[i])
			} else if stack[len(stack)-1] == '*' {
				stack = append(stack, s[i])
			} else {
				stack = stack[:len(stack)-1]
			}
		} else {
			stack = append(stack, s[i])
		}
	}
	return string(stack)
}
