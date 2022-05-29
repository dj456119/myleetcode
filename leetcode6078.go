package myleetcode

func rearrangeCharacters(s string, target string) int {
	cmap := make(map[byte]int)
	for i := range target {
		cmap[target[i]]++
	}
	qmap := make(map[byte]int)
	for i := range s {
		qmap[s[i]]++
	}
	min := 9999999999999
	for k, v := range cmap {
		z := qmap[k]
		if z/v < min {
			min = z / v
		}
	}
	return min
}
