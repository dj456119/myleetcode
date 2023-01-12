package myleetcode

func evaluate(s string, knowledge [][]string) string {
	z := make(map[string]string)
	for i := range knowledge {
		z[knowledge[i][0]] = knowledge[i][1]
	}
	i := 0
	result := []byte{}
	flag := false
	ti := 0
	for ; ; i++ {
		if i == len(s) {
			break
		}
		if s[i] == '(' {
			ti = i
			flag = true
			continue
		}
		if s[i] == ')' {
			c := s[ti+1 : i]
			q := z[c]
			if q == "" {
				result = append(result, '?')
			} else {
				result = append(result, []byte(q)...)
			}

			flag = false
			continue
		}
		if flag {
			continue
		}
		result = append(result, s[i])
	}
	return string(result)
}
