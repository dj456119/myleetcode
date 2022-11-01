package myleetcode

func oddString(words []string) string {
	r := [][]int{}
	for i := range words {
		z := []int{}
		for j := range words[i] {
			if j == 0 {
				continue
			}
			z = append(z, int(words[i][j])-int(words[i][j-1]))
		}
		r = append(r, z)
	}
	for i := range r {
		count := 0
		for j := range r {
			if i == j {
				continue
			}
			for m := range r[i] {
				if r[i][m] != r[j][m] {
					count++
					break
				}
			}
		}
		if count > 1 {
			return words[i]
		}
	}
	return ""
}
