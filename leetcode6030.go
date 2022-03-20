package myleetcode

func longestRepeating(s string, queryCharacters string, queryIndices []int) []int {
	result := []int{}
	ss := []byte(s)
	for i := range queryIndices {
		z := queryCharacters[i]
		pos := queryIndices[i]
		ss[pos] = z
		m := findmax111(ss)
		result = append(result, m)
	}
	return result
}

func findmax111(s []byte) int {
	//dpar1 := make([]int, len(s))
	max := -1
	for i := range s {
		count := 0
		z := s[i]
		for j := i; j < len(s); j++ {
			if s[j] == z {
				count++
			}
			break
		}
		if count > max {
			max = count
		}
		//dpar1[i] = count
	}
	return max
}
